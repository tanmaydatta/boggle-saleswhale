package api

import (
	"github.com/gorilla/mux"
	"github.com/tanmaydatta/boggle/internal"
	"net/http"
)


func SetUp(router *mux.Router) {
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
