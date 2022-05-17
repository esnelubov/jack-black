// Package http_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package http_api

import (
	"time"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for GameAction.
const (
	GameActionBet GameAction = "bet"

	GameActionDoubleDown GameAction = "double_down"

	GameActionEnter GameAction = "enter"

	GameActionHit GameAction = "hit"

	GameActionLose GameAction = "lose"

	GameActionStand GameAction = "stand"

	GameActionTie GameAction = "tie"

	GameActionWin GameAction = "win"
)

// Defines values for StatsRecordResult.
const (
	StatsRecordResultLose StatsRecordResult = "lose"

	StatsRecordResultTie StatsRecordResult = "tie"

	StatsRecordResultWin StatsRecordResult = "win"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// GameAction defines model for GameAction.
type GameAction string

// GameMakeActionPayload defines model for GameMakeActionPayload.
type GameMakeActionPayload struct {
	Action   GameAction `json:"action"`
	ArgsJson *string    `json:"args_json,omitempty"`
	Login    string     `json:"login"`
	SerialId int64      `json:"serial_id"`
}

// GameState defines model for GameState.
type GameState struct {
	AllowedActions  []GameAction `json:"allowed_actions"`
	DescriptionJson string       `json:"description_json"`
	SerialId        int64        `json:"serial_id"`
}

// Player defines model for Player.
type Player struct {
	Balance int64 `json:"balance"`
}

// PlayerCreatePayload defines model for PlayerCreatePayload.
type PlayerCreatePayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Stats defines model for Stats.
type Stats struct {
	History    []StatsRecord `json:"history"`
	TotalLoses int64         `json:"total_loses"`
	TotalWins  int64         `json:"total_wins"`
}

// StatsRecord defines model for StatsRecord.
type StatsRecord struct {
	Balance int64             `json:"balance"`
	Result  StatsRecordResult `json:"result"`
	Time    time.Time         `json:"time"`
}

// StatsRecordResult defines model for StatsRecord.Result.
type StatsRecordResult string

// GameMakeActionJSONBody defines parameters for GameMakeAction.
type GameMakeActionJSONBody GameMakeActionPayload

// GameGetStateParams defines parameters for GameGetState.
type GameGetStateParams struct {
	Login string `json:"login"`
}

// PlayerGetParams defines parameters for PlayerGet.
type PlayerGetParams struct {
	Login string `json:"login"`
}

// PlayerCreateJSONBody defines parameters for PlayerCreate.
type PlayerCreateJSONBody PlayerCreatePayload

// PlayerStatsParams defines parameters for PlayerStats.
type PlayerStatsParams struct {
	Login string `json:"login"`
}

// GameMakeActionJSONRequestBody defines body for GameMakeAction for application/json ContentType.
type GameMakeActionJSONRequestBody GameMakeActionJSONBody

// PlayerCreateJSONRequestBody defines body for PlayerCreate for application/json ContentType.
type PlayerCreateJSONRequestBody PlayerCreateJSONBody