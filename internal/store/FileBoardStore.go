package store

import (
	"bufio"
	"errors"
	"github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tanmaydatta/boggle/internal"
	"github.com/tanmaydatta/boggle/internal/models"
	"log"
	"math/rand"
	"os"
	"strings"
)

const maxBoardSize = 8

var boards map[int][]*models.Board // array so that in future we can have multiple boards

type FileBoardStore struct {}

type boardArray struct {
	board [maxBoardSize][maxBoardSize]string
	vis [maxBoardSize][maxBoardSize]bool
}

func (FileBoardStore) GetBoard(size int) *models.Board {
	l := len(boards[size])
	if l == 0 {
		return nil
	}
	return boards[size][rand.Intn(l)]
}

func (FileBoardStore) createBoard(text string) (b *models.Board, err error) {
	splitted := strings.Split(text, ",")
	l := len(splitted)
	valid, sq := internal.IsSquare(l)
	if !valid {
		err = errors.New("not valid board config")
		return
	}
	var barray [maxBoardSize][maxBoardSize]string
	var chars []string
	for in := range splitted {
		t := strings.TrimSpace(splitted[in])
		if len(t) != 1 {
			err = errors.New("not valid board config")
			return
		}
		barray[in/sq][in%sq] = t
		chars = append(chars, t)
	}
	b = &models.Board{Strings: mapset.NewSet(), Chars: chars, Size: sq}

	generateAllStrings(&boardArray{board: barray}, sq, sq, b)
	return
}

func generateAllStringsUtil(b *boardArray, r, c, x, y int, s string, board *models.Board) {
	if x < 0 || x >= r || y < 0 || y >= c {
		return
	}
	if b.vis[x][y] {
		return
	}
	b.vis[x][y] = true
	toAppend := b.board[x][y]
	if toAppend == "*" {
		toAppend = "."
	}
	s = s + toAppend
	board.Strings.Add(s)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			generateAllStringsUtil(b, r, c, x+i, y+j, s, board)
		}
	}
}

func generateAllStrings(b *boardArray, r, c int, board *models.Board) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			b.vis = [maxBoardSize][maxBoardSize]bool{}
			generateAllStringsUtil(b, r, c, i, j, "", board)
		}
	}
}

func SetupMemBoardStore() FileBoardStore {
	store := &FileBoardStore{}
	boards = make(map[int][]*models.Board)
	file, err := os.Open(viper.GetString("board_path"))
	if err != nil {
		logrus.Fatal("Error opening board file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b, e := store.createBoard(scanner.Text())
		if e != nil {
			continue
		}
		boards[b.Size] = append(boards[b.Size], b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(boards) == 0 {
		log.Fatal("no board configuration found in specified board file")
	}
	return FileBoardStore{}
}