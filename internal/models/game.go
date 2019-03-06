package models

import (
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
	"time"
)

type Game struct {
	Id GameId
	Brd *Board
	StartedAtInSec int32
	DurationInSec int32
	Words mapset.Set
	Attempts mapset.Set
}

type GameId struct {
	xid.ID
}

func (game *Game) IsLive() bool {
	return game.StartedAtInSec + game.DurationInSec > int32(time.Now().Unix())
}


func NewGameId(Id xid.ID) GameId {
	return GameId{ Id}
}