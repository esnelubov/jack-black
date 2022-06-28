package query

import (
	"context"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/history"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/common/generic_errors"
	pm "jackBlack/internal/domain/player"
	sm "jackBlack/internal/domain/player/stats"
	"jackBlack/internal/ports/http_api"
)

type PlayerStatsHandler struct {
	entryRepo   *entry.Repository
	playerRepo  *player.Repository
	historyRepo *history.Repository
}

func NewPlayerStatsHandler(entryRepo *entry.Repository, playerRepo *player.Repository, historyRepo *history.Repository) *PlayerStatsHandler {
	return &PlayerStatsHandler{entryRepo: entryRepo, playerRepo: playerRepo, historyRepo: historyRepo}
}

func (h *PlayerStatsHandler) Handle(ctx context.Context, password string, body *http_api.PlayerStatsParams) (stats *http_api.Stats, err error) {
	var (
		playerAuthorized bool
		playerModel      *pm.Model
		statsModel       *sm.Model
	)

	if playerAuthorized, err = h.entryRepo.PlayerAuthorized(body.Login, password); err != nil {
		return nil, err
	}

	if !playerAuthorized {
		return nil, generic_errors.ErrAuthorization
	}

	if playerModel, err = h.playerRepo.Load(body.Login); err != nil {
		return nil, err
	}

	if statsModel, err = h.historyRepo.Load(playerModel.ID); err != nil {
		return nil, err
	}

	return makeStatsResponseFrom(statsModel), nil
}

func makeStatsResponseFrom(statsModel *sm.Model) (stats *http_api.Stats) {
	stats = &http_api.Stats{
		TotalWins:  statsModel.TotalWins,
		TotalLoses: statsModel.TotalLoses,
		History:    []http_api.StatsRecord{},
	}

	for _, historyRecord := range statsModel.GamesHistory {
		statsRecord := http_api.StatsRecord{
			Balance: int64(*historyRecord.Balance),
			Result:  http_api.StatsRecordResult(historyRecord.Result),
			Time:    historyRecord.Time,
		}

		stats.History = append(stats.History, statsRecord)
	}

	return
}
