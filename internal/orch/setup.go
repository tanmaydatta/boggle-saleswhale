package orch

import "github.com/tanmaydatta/boggle/internal/service"

var boggleService service.BoggleService
var userService service.UserService
var judge service.Judge

func Setup(bs service.BoggleService, us service.UserService, j service.Judge) {
	boggleService = bs
	userService = us
	judge = j
}

