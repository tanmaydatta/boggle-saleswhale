package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal"
	"net/http"
)

func GetGame(router *mux.Router) {
	router.HandleFunc("/{gameId}", internal.Make(
		func(req *http.Request) internal.Response {
			gameId := mux.Vars(req)["gameId"]
			if gameId == "" {
				logrus.Error("invalid game id")
				return internal.BadRequest("invalid game id")
			}
			id, err := boggleService.GetIdFromString(gameId)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			game := boggleService.GetGame(id)
			resp := &GameResp{}
			resp.FromGame(game)
			return internal.Response{
				Code:    http.StatusOK,
				Payload: resp,
			}
		},
	)).Methods(http.MethodGet)
}

func GetScore(router *mux.Router) {
	router.HandleFunc("/{gameId}/score", internal.Make(
		func(req *http.Request) internal.Response {
			gameId := mux.Vars(req)["gameId"]
			if gameId == "" {
				logrus.Error("invalid game id")
				return internal.BadRequest("invalid game id")
			}
			id, err := boggleService.GetIdFromString(gameId)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			game := boggleService.GetGame(id)
			resp := &ScoreResp{judge.GetScore(game)}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: resp,
			}
		},
	)).Methods(http.MethodGet)
}

func SetupBoggleApi(router *mux.Router) {
	s := router.PathPrefix("/game").Subrouter()
	GetGame(s)
	GetScore(s)
}
