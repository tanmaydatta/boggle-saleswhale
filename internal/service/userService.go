package service

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
	"github.com/tanmaydatta/boggle/internal/models"
	"github.com/tanmaydatta/boggle/internal/store/interfaces"
)

type UserService struct {
	Store  interfaces.UserStore
	LiveSt interfaces.LiveGameStore
}

func (u *UserService) StartGameWithNewUser(name string, game *models.Game) (user *models.User, err error) {
	user = u.Store.CreateUser(name)
	err = u.Store.StartGame(user.Id, game.Id)
	if err == nil {
		u.LiveSt.SetUserPlaying(user.Id, game)
	}
	return
}

func (u *UserService) StartGameWithUser(id models.UserId, game *models.Game) (userId models.UserId, err error) {
	if p := u.LiveSt.IsUserPlaying(id); p {
		err = errors.New("user is already playing another game")
		return
	}
	userId = id
	err = u.Store.StartGame(id, game.Id)
	if err == nil {
		u.LiveSt.SetUserPlaying(id, game)
	}
	return
}

func (u *UserService) GetGameIdsOfUser(id models.UserId) (ids mapset.Set){
	ids = u.Store.GetGameIdsOfUser(id)
	return
}

func (u *UserService) PlayMove(id models.UserId, word string) (err error){
	game := u.LiveSt.GetCurrentGameOfUser(id)
	if game == nil || !game.IsLive(){
		err = errors.New("user is not playing any game currently")
		return
	}
	if !game.Brd.Contains(word) {
		err = errors.New(fmt.Sprintf("%s is not present in board", word))
		return
	}
	if game.Attempts.Contains(word) {
		err = errors.New(fmt.Sprintf("already tried %s", word))
		return
	}
	game.Attempts.Add(word)
	return
}

func (u *UserService) AddCorrectMove(id models.UserId, word string) (err error) {
	game := u.LiveSt.GetCurrentGameOfUser(id)
	if game == nil || !game.IsLive(){
		err = errors.New("user is not playing any game currently")
	}
	if game.Words.Contains(word) {
		err = errors.New(fmt.Sprintf("already tried %s", word))
	}
	game.Words.Add(word)
	return
}

func (u *UserService) GetIdFromString(s string) (models.UserId, error) {
	id, e := xid.FromString(s)
	userId := models.NewUserId(id)
	if !u.Store.IsValidUser(userId) {
		return userId, errors.New("invalid userid")
	}
	return userId, e
}

func (u *UserService) GetUser(id models.UserId) * models.User {
	return u.Store.GetUser(id)
}