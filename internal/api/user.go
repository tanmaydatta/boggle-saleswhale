package api

import (
	"encoding/json"
	"github.com/creasty/defaults"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal"
	"net/http"
)

func StartGameNewUser(router *mux.Router) {
	router.HandleFunc("/new/start", internal.Make(
		func(req *http.Request) internal.Response {
			body := internal.GetBytesFromReq(req.Body)
			gameReq := &NewUserGameReq{}
			err := gameReq.DecodeRawAndCreate(body)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			game, err := boggleService.CreateGame(gameReq.DurationInSec, gameReq.Size)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			user, err := userService.StartGameWithNewUser(gameReq.Name, game)
			if err != nil {
				logrus.WithError(err)
				return internal.InternalServerError(err.Error())
			}
			resp := NewUserGameResp{UserId: user.Id, GameId: game.Id, Board: game.Brd.Get2DArray()}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: resp,
			}
		},
	)).Methods(http.MethodPost)
}


func StartGameOldUser(router *mux.Router) {
	router.HandleFunc("/{userId}/start", internal.Make(
		func(req *http.Request) internal.Response {
			userId := mux.Vars(req)["userId"]
			if userId == "" {
				logrus.Error("invalid user id")
				return internal.BadRequest("invalid user id")
			}
			body := internal.GetBytesFromReq(req.Body)
			gameReq := &OldUserGameReq{}
			err := json.Unmarshal(body, gameReq)
			if err != nil {
				logrus.
					WithError(err).
					WithField("raw", body).
					Warn("Could not decode to req req body")
				return internal.BadRequest(err.Error())
			}
			err = defaults.Set(gameReq)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			id, err := userService.GetIdFromString(userId)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			game, err := boggleService.CreateGame(gameReq.DurationInSec, gameReq.Size)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			id , err = userService.StartGameWithUser(id, game)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			resp := NewUserGameResp{UserId: id, GameId: game.Id, Board: game.Brd.Get2DArray()}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: resp,
			}
		},
	)).Methods(http.MethodPost)
}

func PlayMove(router *mux.Router) {
	router.HandleFunc("/{userId}/move", internal.Make(
		func(req *http.Request) internal.Response {
			userId := mux.Vars(req)["userId"]
			if userId == "" {
				logrus.Error("invalid user id")
				return internal.BadRequest("invalid user id")
			}
			body := internal.GetBytesFromReq(req.Body)
			moveReq := &PlayMoveReq{}
			err := json.Unmarshal(body, moveReq)
			if err != nil {
				logrus.
					WithError(err).
					WithField("raw", body).
					Warn("Could not decode to req req body")
				return internal.BadRequest(err.Error())
			}
			if moveReq.Word == "" {
				logrus.Error("empty word provided")
				return internal.BadRequest("empty word provided")
			}
			id, err := userService.GetIdFromString(userId)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			_ = userService.PlayMove(id, moveReq.Word)
			resp := &PlayMoveResp{}
			if judge.IsCorrectWord(moveReq.Word) {
				_ = userService.AddCorrectMove(id, moveReq.Word)
				resp.Correct = true
				resp.Score = judge.GetScoreOfWord(moveReq.Word)
			}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: resp,
			}
		},
	)).Methods(http.MethodPost)
}

func GetUser(router *mux.Router) {
	router.HandleFunc("/{userId}", internal.Make(
		func(req *http.Request) internal.Response {
			userId := mux.Vars(req)["userId"]
			if userId == "" {
				logrus.Error("invalid user id")
				return internal.BadRequest("invalid user id")
			}
			id, err := userService.GetIdFromString(userId)
			if err != nil {
				logrus.WithError(err)
				return internal.BadRequest(err.Error())
			}
			user := userService.GetUser(id)
			if user == nil {
				logrus.Error("invalid user id")
				return internal.BadRequest("invalid user id")
			}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: user,
			}
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