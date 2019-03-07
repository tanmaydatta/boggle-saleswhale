package validation

import "github.com/tanmaydatta/boggle/internal/service"

var boggleService service.BoggleService
var userService service.UserService

func Setup(bs service.BoggleService, us service.UserService) {
	boggleService = bs
	userService = us
}
