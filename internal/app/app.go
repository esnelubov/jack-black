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
	"jackBlack/internal/app/query"
	"os"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	PlayerCreate   *command.PlayerCreateHandler
	GameMakeAction *command.GameMakeActionHandler
}

type Queries struct {
	Player      *query.PlayerHandler
	PlayerStats *query.PlayerStatsHandler
	GameState   *query.GameStateHandler
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
			PlayerCreate:   command.NewPlayerCreateHandler(entryRepository, playerRepository),
			GameMakeAction: command.NewGameMakeActionHandler(entryRepository, playerRepository, sessionRepository, historyRepository),
		},
		Queries: Queries{
			Player:      query.NewPlayerHandler(entryRepository, playerRepository),
			PlayerStats: query.NewPlayerStatsHandler(entryRepository, playerRepository, historyRepository),
			GameState:   query.NewGameStateHandler(entryRepository, playerRepository, sessionRepository),
		},
	}
}
