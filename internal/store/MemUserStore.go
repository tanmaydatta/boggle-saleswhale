package store

import (
	"errors"
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
	"github.com/tanmaydatta/boggle/internal/models"
	"time"
)

var userStoreMap map[models.UserId]*models.User

type MemUserStore struct {}

func (MemUserStore) IsValidUser(id models.UserId) bool {
	return userStoreMap[id] != nil
}

func (MemUserStore) GetUser(id models.UserId) *models.User {
	return userStoreMap[id]
}

func (MemUserStore) StartGame(id models.UserId, gameId models.GameId) (err error) {
	user := userStoreMap[id]
	if user == nil {
		err = errors.New("invalid userid")
		return
	}
	err = nil
	added := user.Games.Add(gameId)
	if !added {
		err = errors.New("user is already playing this game")
	}
	return
}

func (MemUserStore) CreateUser(name string) (user *models.User) {
	user = &models.User{Id: models.NewUserId(xid.NewWithTime(time.Now())), Name: name, Games: mapset.NewSet()}
	userStoreMap[user.Id] = user
	return
}

func (MemUserStore) GetGameIdsOfUser(id models.UserId) mapset.Set {
	user := userStoreMap[id]
	if user == nil {
		return mapset.NewSet()
	}
	return user.Games
}


func SetupMemUserStore() MemUserStore {
	userStoreMap = make(map[models.UserId]*models.User)
	return MemUserStore{}
}