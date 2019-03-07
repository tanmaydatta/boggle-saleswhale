package validation

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/tanmaydatta/boggle/internal/models"
	"net/http"
)

func ValidateGetGameReq(req *http.Request) (id models.GameId, err error) {
	gameId := mux.Vars(req)["gameId"]
	return ValidateGameId(gameId)
}

func ValidateGetScoreReq(req *http.Request) (id models.GameId, err error) {
	gameId := mux.Vars(req)["gameId"]
	return ValidateGameId(gameId)
}

func ValidateGameId(gameId string) (id models.GameId, err error)  {
	if gameId == "" {
		err = errors.New("invalid game id")
	}
	id, err = boggleService.GetIdFromString(gameId)
	if err != nil {
		return
	}
	if boggleService.GetGame(id) == nil {
		return id, errors.New("invalid game id")
	}
	return
}
