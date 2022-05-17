package app

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jackBlack/internal/adapters/gorm/repositories/entry"
	"jackBlack/internal/adapters/gorm/repositories/history"
	"jackBlack/internal/adapters/gorm/repositories/player"
	"jackBlack/internal/adapters/gorm/repositories/session"
	"jackBlack/internal/app/command"
	"os"
)

type Application struct {
	Commands Commands
}

type Commands struct {
	PlayerGet *command.PlayerGetHandler

	PlayerCreate *command.PlayerCreateHandler

	PlayerStats *command.PlayerStatsHandler

	GameMakeAction *command.GameMakeActionHandler

	GameGetState *command.GameGetStateHandler
}

func NewApplication(ctx context.Context) *Application {
	var err error

	dsn := os.Getenv("BJ_DSN")

	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres"
	}

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
		})
	if err != nil {
		panic(err)
	}

	entryRepository := entry.New(db)
	historyRepository := history.New(db)
	playerRepository := player.New(db)
	sessionRepository := session.New(db)

	return &Application{
		Commands: Commands{
			PlayerGet:      command.NewPlayerGetHandler(entryRepository, playerRepository),
			PlayerCreate:   command.NewPlayerCreateHandler(entryRepository, playerRepository),
			PlayerStats:    command.NewPlayerStatsHandler(entryRepository, playerRepository, historyRepository),
			GameMakeAction: command.NewGameMakeActionHandler(entryRepository, playerRepository, sessionRepository, historyRepository),
			GameGetState:   command.NewGameGetStateHandler(entryRepository, playerRepository, sessionRepository),
		},
	}
}
