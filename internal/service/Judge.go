package service

import (
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store"
)

type Judge struct {
	Dictionary store.DictionaryStore
}

func (j *Judge) IsCorrectWord(word string) bool {
	return j.Dictionary.Contains(word)
}

func (j *Judge) GetScore(game *models.Game) (score int32) {
	score = 0
	for w := range game.Words.Iter() {
		word := w.(string)
		l := len(word)
		if l <= 4 {
			score++
		} else if l == 5 {
			score += 2
		} else if l == 6 {
			score += 3
		} else if l == 7 {
			score += 5
		} else {
			score+=11
		}
	}
	return
}
