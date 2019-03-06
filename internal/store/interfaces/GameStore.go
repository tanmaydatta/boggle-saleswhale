package interfaces

import (
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store"
)

type GameStore interface {
	GetGame(id models.GameId) *models.Game
	IsGameLive(id models.GameId) bool
	CreateGame(durationInSec int32, b *models.Board) *models.Game
	AddAttempt(id models.GameId, attempt string)
	AddCorrectWord(id models.GameId, word string)
	IsValidGame(id models.GameId) bool
}

func SetupGameStore() GameStore {
	return store.SetupMemGameStore()
}
