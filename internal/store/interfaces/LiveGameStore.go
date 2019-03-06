package interfaces

import (
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store"
)

type LiveGameStore interface {
	IsUserPlaying(id models.UserId) bool
	SetUserPlaying(id models.UserId, g *models.Game)
	GetCurrentGameOfUser(id models.UserId) *models.Game
}

func SetupLiveGameStore() LiveGameStore {
	return store.SetupMemLiveGameStore()
}
