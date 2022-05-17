package bet

import (
	"fmt"
	"jackBlack/internal/common/set_of_strings"
	"jackBlack/internal/domain/game"
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/money"
)

type Action string

const BetArg = "bet"

func (this Action) AllowedFor(session action.GameData) bool {
	return *session.Player().Balance >= game.MinimalBet
}

func (this Action) Perform(session action.GameData, args map[string]interface{}) (err error) {
	betArg, ok := args[BetArg]
	if !ok {
		return action.MakeErrArgNorProvided(BetArg)
	}

	bet, ok := betArg.(float64) // Defaults to float in json
	if !ok {
		return action.MakeErrArgParse(betArg, "float64")
	}

	if bet < game.MinimalBet {
		return fmt.Errorf("bet (%.0f) is lower than the minimal bet (%d)", bet, game.MinimalBet)
	}

	if err = session.Player().Balance.Transfer(*money.New(int64(bet)), session.Bet()); err != nil {
		return fmt.Errorf("bet (%.0f) is higher than the player's balance (%d)", bet, session.Player().Balance)
	}

	if err = session.Deck().DealIntoEach(2, session.PlayerHand(), session.DealerHand()); err != nil {
		return err
	}

	return
}

func (this Action) DescribeState(session action.GameData) (description map[string]interface{}) {
	return map[string]interface{}{
		"balance":     session.Player().Balance,
		"bet":         session.Bet(),
		"player_hand": session.PlayerHand(),
		"dealer_card": game.ShowDealersFirstCard(session.DealerHand()),
	}
}

func (this Action) NextActions(session action.GameData) (allowedActions set_of_strings.Type) {
	return action.NextActionsCommon(session)
}

func (from Action) Transit(to string, session action.GameData, args map[string]interface{}) (err error) {
	return action.TransitCommon(from, to, session, args)
}

func (this Action) Name() string {
	return string(this)
}
