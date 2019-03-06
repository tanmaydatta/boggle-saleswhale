package models

import (
	"github.com/deckarep/golang-set"
	"github.com/rs/xid"
)

type User struct {
	Id UserId			`json:"id"`
	Name string			`json:"name"`
	Games mapset.Set	`json:"games"`
}

type UserId struct {
	xid.ID
}

func NewUserId(Id xid.ID) UserId {
	return UserId{Id}
}