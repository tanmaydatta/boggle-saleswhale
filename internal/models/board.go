package models

import (
	"github.com/deckarep/golang-set"
	"regexp"
)

type Board struct {
	Strings mapset.Set
	Chars []string
	Size int
}

func (b *Board) Contains(word string) bool {
	for str := range b.Strings.Iter() {
		match, _ := regexp.MatchString(str.(string), word)
		if match {
			return match
		}
	}
	return false
}

func (b *Board) Get2DArray() [][]string {
	res := make([][]string, b.Size)
	for i := 0; i < b.Size; i++ {
		res[i] = make([]string, b.Size)
	}
	for i := 0; i < len(b.Chars); i++ {
		res[i/b.Size][i%b.Size] = b.Chars[i]
	}
	return res
}
