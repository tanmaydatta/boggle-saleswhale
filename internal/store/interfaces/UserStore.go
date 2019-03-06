package interfaces

import (
	"github.com/deckarep/golang-set"
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store"
)

type UserStore interface {
	StartGame(id models.UserId, gameId models.GameId) error
	CreateUser(name string) *models.User
	GetGameIdsOfUser(id models.UserId) mapset.Set
	GetUser(id models.UserId) *models.User
	IsValidUser(id models.UserId) bool
}

func SetupUserStore() UserStore {
	return store.SetupMemUserStore()
}