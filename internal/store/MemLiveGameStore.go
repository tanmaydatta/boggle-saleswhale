package store

import (
	"github.com/tanmaydatta/boggle/internal/models"
)

type MemLiveGameStore struct {
	UserMap map[models.UserId]*models.Game
}

func (st MemLiveGameStore) SetUserPlaying(id models.UserId, g *models.Game) {
	st.UserMap[id] = g
}

func (st MemLiveGameStore) GetCurrentGameOfUser(id models.UserId) *models.Game {
	return st.UserMap[id]
}

func (st MemLiveGameStore) IsUserPlaying(id models.UserId) bool {
	game  := st.UserMap[id]
	if game == nil {
		return false
	}
	if game.IsLive() {
		return true
	}
	delete(st.UserMap, id)
	return false
}

func SetupMemLiveGameStore() MemLiveGameStore {
	userMap := make(map[models.UserId]*models.Game)
	return MemLiveGameStore{UserMap: userMap}
}