package models

import (
	"encoding/json"
	"errors"
	"github.com/creasty/defaults"
	"github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
	"github.com/tanmaydatta/boggle/internal/models"
)

type NewUserGameReq struct {
	Name string				`json:"name"`
	DurationInSec int32 	`json:"duration" default:"10"`
	Size int				`json:"size" default:"4"`
}

type OldUserGameReq struct {
	DurationInSec int32 	`json:"duration" default:"60"`
	Size int				`json:"size" default:"4"`
}

type PlayMoveReq struct {
	Word string `json:"word"`
}

type PlayMoveResp struct {
	Correct bool `json:"correct"`
	Score int32	 `json:"score"`
}

type NewUserGameResp struct {
	UserId models.UserId	`json:"userid"`
	GameId models.GameId 	`json:"gameid"`
	Board [][]string		`json:"board"`
}

type GameResp struct {
	Id models.GameId		`json:"id"`
	StartedAtInSec int32	`json:"startedAtInSec"`
	DurationInSec int32		`json:"durationInSec"`
	Words mapset.Set		`json:"words"`
	Attempts mapset.Set		`json:"Attempts"`
	Board [][]string		`json:"board"`
}

type ScoreResp struct {
	Score int32 `json:"score"`
}

func (req *NewUserGameReq) DecodeRawAndCreate(raw []byte) (er error) {
	err := json.Unmarshal(raw, req)
	if err != nil {
		logrus.
			WithError(err).
			WithField("raw", raw).
			Warn("Could not decode to req req body")
		return errors.New(err.Error())
	}
	er = defaults.Set(req)
	return
}

func (g *GameResp) FromGame(game *models.Game) {
	g.Id = game.Id
	g.StartedAtInSec = game.StartedAtInSec
	g.DurationInSec = game.DurationInSec
	g.Words = game.Words
	g.Attempts = game.Attempts
	g.Board = game.Brd.Get2DArray()
}