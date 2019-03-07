package orch

import (
	"github.com/tanmaydatta/boggle/internal"
	m "github.com/tanmaydatta/boggle/internal/api/models"
	"github.com/tanmaydatta/boggle/internal/models"
	"net/http"
)

func OrchGetGame(id models.GameId) internal.Response {
	game := boggleService.GetGame(id)
	resp := &m.GameResp{}
	resp.FromGame(game)
	return internal.Response{
		Code:    http.StatusOK,
		Payload: resp,
	}
}

func OrchGetScore(id models.GameId) internal.Response {
	game := boggleService.GetGame(id)
	resp := &m.ScoreResp{judge.GetScore(game)}
	return internal.Response{
		Code:    http.StatusOK,
		Payload: resp,
	}
}