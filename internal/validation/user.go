package validation

import (
	"encoding/json"
	"errors"
	"github.com/creasty/defaults"
	"github.com/gorilla/mux"
	"github.com/tanmaydatta/boggle/internal"
	m "github.com/tanmaydatta/boggle/internal/api/models"
	"github.com/tanmaydatta/boggle/internal/models"
	"net/http"
)

func ValidateStartGameOldUserReq(req *http.Request) (id models.UserId, gameReq *m.OldUserGameReq, err error) {
	userId := mux.Vars(req)["userId"]
	id, err = ValidateUserId(userId)
	if err != nil {
		return
	}
	body := internal.GetBytesFromReq(req.Body)
	gameReq = &m.OldUserGameReq{}
	err = json.Unmarshal(body, gameReq)
	if err != nil {
		return
	}
	err = defaults.Set(gameReq)
	return
}

func ValidateStartGameWithNewUserReq(req *http.Request) (gameReq *m.NewUserGameReq, err error) {
	body := internal.GetBytesFromReq(req.Body)
	gameReq = &m.NewUserGameReq{}
	err = gameReq.DecodeRawAndCreate(body)
	return
}

func ValidatePlayMoveReq(req *http.Request) (id models.UserId, moveReq *m.PlayMoveReq, err error) {
	userId := mux.Vars(req)["userId"]
	id, err = ValidateUserId(userId)
	if err != nil {
		return
	}
	body := internal.GetBytesFromReq(req.Body)
	moveReq = &m.PlayMoveReq{}
	err = json.Unmarshal(body, moveReq)
	if err != nil {
		return
	}
	if moveReq.Word == "" {
		err =  errors.New("empty word provided")
	}
	return
}

func ValidateGetUserReq(req *http.Request) (id models.UserId, err error) {
	userId := mux.Vars(req)["userId"]
	return ValidateUserId(userId)
}

func ValidateUserId(userId string) (id models.UserId, err error)  {
	if userId == "" {
		err = errors.New("invalid user id")
	}
	id, err = userService.GetIdFromString(userId)
	if err != nil {
		return
	}
	if userService.GetUser(id) == nil {
		err = errors.New("invalid user id")
	}
	return
}

