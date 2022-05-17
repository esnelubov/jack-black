package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"jackBlack/internal/common/client/http_api"
	"jackBlack/internal/domain/game/deck"
	"jackBlack/internal/domain/game/deck/card"
	"jackBlack/internal/domain/game/money"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		client *http_api.ClientWithResponses
		err    error
		sc     = bufio.NewScanner(os.Stdin)
	)

	addr := getFromUser(sc, "Please input server address", "http://127.0.0.1:3434")
	addr = strings.TrimSuffix(addr, "/")
	addr = fmt.Sprintf("%s/api", addr)

	if client, err = http_api.NewClientWithResponses(addr); err != nil {
		panic(err)
	}

	login := getFromUser(sc, "Please input your login", "test")
	password := getFromUser(sc, "Please input your password", "test")

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		resp, err := client.PlayerGetWithResponse(ctx, &http_api.PlayerGetParams{Login: login}, makeAuthFunc(password))
		if err != nil {
			panic(err)
		}
		if resp.StatusCode() == http.StatusUnauthorized {
			password = getFromUser(sc, "Password is incorrect. Please input it again", "test")
		}
		if resp.StatusCode() == http.StatusNotFound {
			resp, err := client.PlayerCreateWithResponse(ctx, http_api.PlayerCreateJSONRequestBody{Login: login, Password: password})
			if err != nil {
				panic(err)
			}

			fmt.Println("Player was not found. Created automatically.")

			if resp.StatusCode() == http.StatusOK {
				break
			}
		}
		if resp.StatusCode() == http.StatusOK {
			break
		}
	}

	startActionsLoop(sc, client, login, password)
}

func getFromUser(scanner *bufio.Scanner, prompt string, def string) string {
	fmt.Printf("%s (default: %s)\n", prompt, def)
	scanner.Scan()
	result := scanner.Text()
	if result == "" {
		result = def
	}

	return result
}

func makeAuthFunc(password string) http_api.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-API-KEY", password)
		return nil
	}
}

func startActionsLoop(sc *bufio.Scanner, client *http_api.ClientWithResponses, login, password string) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		state, err := client.GameGetStateWithResponse(ctx, &http_api.GameGetStateParams{Login: login}, makeAuthFunc(password))
		if err != nil {
			panic(err)
		}
		fmt.Println("===============================")
		printDescription(state.JSON200.DescriptionJson)
		fmt.Println("-------------------------------")
		printAllowedActions(state.JSON200.AllowedActions)
		serialID := state.JSON200.SerialId
		serialID++
		makeAction(sc, client, login, password, serialID)
	}
}

type GameStateDescription struct {
	Balance    *money.Model `json:"balance"`
	Bet        *money.Model `json:"bet"`
	PlayerHand *deck.Model  `json:"player_hand"`
	DealerHand *deck.Model  `json:"dealer_hand"`
	DealerCard *card.Model  `json:"dealer_card"`
}

func printDescription(description string) {
	var decoded *GameStateDescription
	if err := json.Unmarshal([]byte(description), &decoded); err != nil {
		panic(err)
	}

	printHand := func(hand *deck.Model) {
		for _, c := range *hand {
			fmt.Printf("    %v\n", c)
		}
	}

	if decoded.Balance != nil {
		fmt.Printf("Your balance: %d\n", *decoded.Balance)
	}

	if decoded.Bet != nil {
		fmt.Printf("Your bet: %d\n", *decoded.Bet)
	}

	if decoded.PlayerHand != nil {
		fmt.Print("Your hand:\n")
		printHand(decoded.PlayerHand)
	}

	if decoded.DealerHand != nil {
		fmt.Print("Dealer's hand:\n")
		printHand(decoded.DealerHand)
	}

	if decoded.DealerCard != nil {
		fmt.Printf("Dealer shows the card: %+v\n", decoded.DealerCard)
	}
}

var ActionToPrompt = map[http_api.GameAction]string{
	http_api.GameActionBet:        "b [amount] - Bet",
	http_api.GameActionDoubleDown: "d - Double Down",
	http_api.GameActionEnter:      "e - Enter New Game",
	http_api.GameActionHit:        "h - Hit",
	http_api.GameActionLose:       "l - Accept lose",
	http_api.GameActionStand:      "s - Stand",
	http_api.GameActionTie:        "t - Accept Tie",
	http_api.GameActionWin:        "w - Collect winning",
}

func printAllowedActions(allowedActions []http_api.GameAction) {
	fmt.Println("Input:")
	fmt.Println("    i - Account info")
	for _, a := range allowedActions {
		fmt.Println("    " + ActionToPrompt[a])
	}
}

var buttonRegexp = regexp.MustCompile(`(?P<button>\S)\s*(?P<args>\S*)`)

var ButtonToAction = map[string]http_api.GameAction{
	"b": http_api.GameActionBet,
	"d": http_api.GameActionDoubleDown,
	"e": http_api.GameActionEnter,
	"h": http_api.GameActionHit,
	"l": http_api.GameActionLose,
	"s": http_api.GameActionStand,
	"t": http_api.GameActionTie,
	"w": http_api.GameActionWin,
}

func makeAction(sc *bufio.Scanner, client *http_api.ClientWithResponses, login, password string, serialID int64) {
	sc.Scan()
	input := sc.Text()
	matches := buttonRegexp.FindStringSubmatch(input)
	if len(matches) == 0 {
		return
	}

	button := matches[1]

	arg := ""
	if len(matches) > 2 {
		arg = matches[2]
	}

	if button == "i" {
		printPlayerInfo(client, login, password)
		return
	}

	action, ok := ButtonToAction[button]
	if !ok {
		fmt.Printf("Unexpected button: %s\n", button)
		return
	}

	argsMap := map[string]interface{}{}

	if action == http_api.GameActionBet {
		if arg == "" {
			fmt.Printf("Bet is not provided\n")
			return
		}

		bet, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Bet must be integer\n")
			return
		}

		argsMap["bet"] = bet
	}

	argsJson, err := json.Marshal(argsMap)
	if err != nil {
		panic(err)
	}

	argsJsonString := string(argsJson)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := client.GameMakeActionWithResponse(ctx, http_api.GameMakeActionJSONRequestBody{
		Action:   action,
		ArgsJson: &argsJsonString,
		Login:    login,
		SerialId: serialID,
	}, makeAuthFunc(password))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("ERROR: %s\n", resp.JSONDefault.Message)
	}
}

func printPlayerInfo(client *http_api.ClientWithResponses, login, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := client.PlayerStatsWithResponse(ctx, &http_api.PlayerStatsParams{Login: login}, makeAuthFunc(password))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total wins: %d\n", resp.JSON200.TotalWins)
	fmt.Printf("Total loses: %d\n", resp.JSON200.TotalLoses)

	if resp.JSON200.History != nil {
		fmt.Println("Games history:")
		for _, h := range resp.JSON200.History {
			fmt.Printf("%s\n", h.Time.Format(time.RFC1123))
			fmt.Printf("    Result: %s\n", h.Result)
			fmt.Printf("    Balance: %d\n", h.Balance)
			fmt.Println()
		}
	}
}
