package api

import (
	"github.com/gorilla/mux"
	"github.com/tanmaydatta/boggle/internal"
	"github.com/tanmaydatta/boggle/internal/service"
	"github.com/tanmaydatta/boggle/internal/store"
	"github.com/tanmaydatta/boggle/internal/store/interfaces"
	"net/http"
)

var boggleService service.BoggleService
var userService service.UserService
var judge service.Judge

func SetUp(router *mux.Router) {
	ust := interfaces.SetupUserStore()
	gst := interfaces.SetupGameStore()
	lgst := interfaces.SetupLiveGameStore()
	bst := interfaces.SetupBoardStore()
	dic := store.SetupDictionaryStore()
	judge = service.Judge{Dictionary: dic}
	userService = service.UserService{Store: ust, LiveSt: lgst}
	boggleService = service.BoggleService{Store: gst, Bst: bst}
	addHealthApi(router)
	SetupUserApi(router)
	SetupBoggleApi(router)
}

func addHealthApi(router *mux.Router) {
	router.HandleFunc("/health", internal.Make(
		func(req *http.Request) internal.Response {
			response := internal.Fields{
				"message": "Server is running",
			}
			return internal.Response{
				Code:    http.StatusOK,
				Payload: response,
			}
		},
	)).Methods(http.MethodGet)
}


// Todo: get score
// Todo: play move: done
// Todo: get game: done