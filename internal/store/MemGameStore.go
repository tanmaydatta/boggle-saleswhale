package store

import (
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
	"github.com/tanmaydatta/boggle/internal/models"
	"time"
)

var gameMap map[models.GameId]*models.Game

type MemGameStore struct {}

func (st MemGameStore) GetGame(id models.GameId) *models.Game {
	return gameMap[id]
}

func (st MemGameStore) IsGameLive(id models.GameId) bool {
	game := st.GetGame(id)
	if game == nil {
		return false
	}
	return game.IsLive()
}

func (st MemGameStore) AddAttempt(id models.GameId, attempt string) {
	game := st.GetGame(id)
	if game == nil {
		return
	}
	game.Attempts.Add(attempt)
}

func (st MemGameStore) AddCorrectWord(id models.GameId, word string) {
	game := st.GetGame(id)
	if game == nil {
		return
	}
	game.Words.Add(word)
}

func (st MemGameStore) CreateGame(durationInSec int32, board *models.Board) *models.Game {
	g := models.Game{
		Id: models.NewGameId(xid.NewWithTime(time.Now())),
		DurationInSec: durationInSec,
		Brd: board,
		StartedAtInSec: int32(time.Now().Unix()),
		Attempts: mapset.NewSet(),
		Words: mapset.NewSet(),
	}
	gameMap[g.Id] = &g
	return &g
}

func (MemGameStore) IsValidGame(id models.GameId) bool {
	return gameMap[id] != nil
}

func SetupMemGameStore() MemGameStore {
	gameMap = make(map[models.GameId]*models.Game)
	return MemGameStore{}
}