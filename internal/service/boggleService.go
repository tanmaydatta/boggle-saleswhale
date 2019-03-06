package service

import (
	"errors"
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store/interfaces"
)

type BoggleService struct {
	Store interfaces.GameStore
	Bst interfaces.BoardStore
}

func (g *BoggleService) GetGame(id models.GameId) *models.Game {
	return g.Store.GetGame(id)
}

func (g *BoggleService) GetGames(ids mapset.Set) (res []*models.Game) {
	games := make(chan *models.Game, ids.Cardinality())
	for id := range ids.Iter() {
		go func() {games <- g.GetGame(id.(models.GameId))}()
	}
	for i := 0; i < ids.Cardinality(); i++ {
		res = append(res, <-games)
	}
	return
}

func (g *BoggleService) CreateGame(duration int32, size int) (*models.Game, error) {
	board := g.Bst.GetBoard(size)
	if board == nil {
		return nil, errors.New("board not found")
	}
	return g.Store.CreateGame(duration, board), nil
}

func (g *BoggleService) GetIdFromString(s string) (models.GameId, error) {
	id, e := xid.FromString(s)
	gameId := models.NewGameId(id)
	if !g.Store.IsValidGame(gameId) {
		return gameId, errors.New("invalid userid")
	}
	return gameId, e
}