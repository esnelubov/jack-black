package actions

import (
	"jackBlack/internal/domain/game/action"
	"jackBlack/internal/domain/game/actions/bet"
	"jackBlack/internal/domain/game/actions/double_down"
	"jackBlack/internal/domain/game/actions/enter"
	"jackBlack/internal/domain/game/actions/hit"
	"jackBlack/internal/domain/game/actions/lose"
	"jackBlack/internal/domain/game/actions/stand"
	"jackBlack/internal/domain/game/actions/tie"
	"jackBlack/internal/domain/game/actions/win"
	"jackBlack/internal/domain/game/constants"
)

const (
	Enter      = enter.Action(constants.ActionEnter)
	Bet        = bet.Action(constants.ActionBet)
	Hit        = hit.Action(constants.ActionHit)
	Stand      = stand.Action(constants.ActionStand)
	DoubleDown = double_down.Action(constants.ActionDoubleDown)
	Win        = win.Action(constants.ActionWin)
	Lose       = lose.Action(constants.ActionLose)
	Tie        = tie.Action(constants.ActionTie)
)

func InitActionsMap() {
	action.WithName = action.ActionMap{
		Enter.Name():      Enter,
		Bet.Name():        Bet,
		Hit.Name():        Hit,
		Stand.Name():      Stand,
		DoubleDown.Name(): DoubleDown,
		Win.Name():        Win,
		Lose.Name():       Lose,
		Tie.Name():        Tie,
	}
}
