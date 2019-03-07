package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal"
	"github.com/tanmaydatta/boggle/internal/orch"
	"github.com/tanmaydatta/boggle/internal/validation"
	"net/http"
)

func StartGameNewUser(router *mux.Router) {
	router.HandleFunc("/new/start", internal.Make(
		func(req *http.Request) internal.Response {
			gameReq, err := validation.ValidateStartGameWithNewUserReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchStartGameNewUser(gameReq)
		},
	)).Methods(http.MethodPost)
}


func StartGameOldUser(router *mux.Router) {
	router.HandleFunc("/{userId}/start", internal.Make(
		func(req *http.Request) internal.Response {
			id, gameReq, err := validation.ValidateStartGameOldUserReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchStartGameOldUser(id, gameReq)
		},
	)).Methods(http.MethodPost)
}

func PlayMove(router *mux.Router) {
	router.HandleFunc("/{userId}/move", internal.Make(
		func(req *http.Request) internal.Response {
			id, moveReq, err := validation.ValidatePlayMoveReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchPlayMove(id, moveReq)
		},
	)).Methods(http.MethodPost)
}

func GetUser(router *mux.Router) {
	router.HandleFunc("/{userId}", internal.Make(
		func(req *http.Request) internal.Response {
			id, err := validation.ValidateGetUserReq(req)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			return orch.OrchGetUser(id)
		},
	)).Methods(http.MethodGet)
}

func SetupUserApi(router *mux.Router) {
	s := router.PathPrefix("/users").Subrouter()
	StartGameNewUser(s)
	GetUser(s)
	StartGameOldUser(s)
	PlayMove(s)
}