package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal"
	"github.com/tanmaydatta/boggle/internal/orch"
	"github.com/tanmaydatta/boggle/internal/validation"
	"net/http"
)

func GetGame(router *mux.Router) {
	router.HandleFunc("/{gameId}", internal.Make(
		func(req *http.Request) internal.Response {
			id, err := validation.ValidateGetGameReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchGetGame(id)
		},
	)).Methods(http.MethodGet)
}

func GetScore(router *mux.Router) {
	router.HandleFunc("/{gameId}/score", internal.Make(
		func(req *http.Request) internal.Response {
			id, err := validation.ValidateGetScoreReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchGetScore(id)
		},
	)).Methods(http.MethodGet)
}

func SetupBoggleApi(router *mux.Router) {
	s := router.PathPrefix("/game").Subrouter()
	GetGame(s)
	GetScore(s)
}
