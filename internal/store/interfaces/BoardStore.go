package interfaces

import (
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store"
)

type BoardStore interface {
	GetBoard(size int) *models.Board
}

func SetupBoardStore() BoardStore {
	return store.SetupMemBoardStore()
}
