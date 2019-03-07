package orch

import (
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal"
	m "github.com/tanmaydatta/boggle/internal/api/models"
	"github.com/tanmaydatta/boggle/internal/models"
	"net/http"
)

func OrchGetUser(id models.UserId) internal.Response {
	user := userService.GetUser(id)
	if user == nil {
		logrus.Error("invalid user id")
		return internal.BadRequest("invalid user id")
	}
	return internal.Response{
		Code:    http.StatusOK,
		Payload: user,
	}
}

func OrchPlayMove(id models.UserId, moveReq *m.PlayMoveReq) internal.Response {
	_ = userService.PlayMove(id, moveReq.Word)
	resp := &m.PlayMoveResp{}
	if judge.IsCorrectWord(moveReq.Word) {
		_ = userService.AddCorrectMove(id, moveReq.Word)
		resp.Correct = true
		resp.Score = judge.GetScoreOfWord(moveReq.Word)
	}
	return internal.Response{
		Code:    http.StatusOK,
		Payload: resp,
	}
}

func OrchStartGameOldUser(id models.UserId, gameReq *m.OldUserGameReq) internal.Response {
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
	resp := m.NewUserGameResp{UserId: id, GameId: game.Id, Board: game.Brd.Get2DArray()}
	return internal.Response{
		Code:    http.StatusOK,
		Payload: resp,
	}
}

func OrchStartGameNewUser(gameReq *m.NewUserGameReq) internal.Response {
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
	resp := m.NewUserGameResp{UserId: user.Id, GameId: game.Id, Board: game.Brd.Get2DArray()}
	return internal.Response{
		Code:    http.StatusOK,
		Payload: resp,
	}
}