package ports

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"jackBlack/internal/app"
	"jackBlack/internal/common/generic_errors"
	"jackBlack/internal/ports/http_api"
	"net/http"
)

type HttpServer struct {
	app *app.Application
}

func NewHttpServer(application *app.Application) *HttpServer {
	return &HttpServer{
		app: application,
	}
}

func (h *HttpServer) GameMakeAction(w http.ResponseWriter, r *http.Request) {
	var (
		password string
		body     http_api.GameMakeActionPayload
		err      error
	)

	if password, err = getPassword(r); err != nil {
		respondWithError(w, r, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, r, err)
		return
	}

	if err = h.app.Commands.GameMakeAction.Handle(r.Context(), password, &body); err != nil {
		respondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpServer) GameGetState(w http.ResponseWriter, r *http.Request, params http_api.GameGetStateParams) {
	var (
		password  string
		gameState *http_api.GameState
		err       error
	)

	if password, err = getPassword(r); err != nil {
		respondWithError(w, r, err)
		return
	}

	if gameState, err = h.app.Commands.GameGetState.Handle(r.Context(), password, &params); err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, gameState)
}

func (h *HttpServer) PlayerGet(w http.ResponseWriter, r *http.Request, params http_api.PlayerGetParams) {
	var (
		password string
		player   *http_api.Player
		err      error
	)

	if password, err = getPassword(r); err != nil {
		respondWithError(w, r, err)
		return
	}

	if player, err = h.app.Commands.PlayerGet.Handle(r.Context(), password, &params); err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, player)
}

func (h *HttpServer) PlayerCreate(w http.ResponseWriter, r *http.Request) {
	var (
		body   http_api.PlayerCreatePayload
		player *http_api.Player
		err    error
	)

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithError(w, r, err)
		return
	}

	if player, err = h.app.Commands.PlayerCreate.Handle(r.Context(), &body); err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, player)
}

func (h *HttpServer) PlayerStats(w http.ResponseWriter, r *http.Request, params http_api.PlayerStatsParams) {
	var (
		password string
		stats    *http_api.Stats
		err      error
	)

	if password, err = getPassword(r); err != nil {
		respondWithError(w, r, err)
		return
	}

	if stats, err = h.app.Commands.PlayerStats.Handle(r.Context(), password, &params); err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, stats)
}

func getPassword(r *http.Request) (password string, err error) {
	password = r.Header.Get("X-API-KEY")
	if password == "" {
		err = generic_errors.ErrAuthorization
	}
	return
}

func respondWithError(w http.ResponseWriter, r *http.Request, err error) {
	bjErr := http_api.Error{
		Message: err.Error(),
	}

	code := http.StatusInternalServerError

	if errors.Is(err, generic_errors.ErrAuthorization) {
		code = http.StatusUnauthorized
	}

	if errors.Is(err, generic_errors.ErrPlayerNotExists) {
		code = http.StatusNotFound
	}

	if errors.Is(err, generic_errors.ErrLoginAlreadyTaken) {
		code = http.StatusConflict
	}

	if errors.Is(err, generic_errors.ErrActionIsNotAllowed) {
		code = http.StatusNotAcceptable
	}

	render.Status(r, code)
	render.JSON(w, r, bjErr)
}
