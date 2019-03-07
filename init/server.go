package server

import (
	"github.com/betacraft/yaag/middleware"
	"github.com/betacraft/yaag/yaag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tanmaydatta/boggle/internal/api"
	"github.com/tanmaydatta/boggle/internal/orch"
	"github.com/tanmaydatta/boggle/internal/service"
	"github.com/tanmaydatta/boggle/internal/store"
	"github.com/tanmaydatta/boggle/internal/store/interfaces"
	"github.com/tanmaydatta/boggle/internal/validation"
	"net/http"
	"time"
)

type Server struct {
	*mux.Router
	Address string
}

func (s Server) SetupComponents() {
	apiMux := s.PathPrefix("/api").Subrouter()
	ust := interfaces.SetupUserStore()
	gst := interfaces.SetupGameStore()
	lgst := interfaces.SetupLiveGameStore()
	bst := interfaces.SetupBoardStore()
	dic := store.SetupDictionaryStore()
	judge := service.Judge{Dictionary: dic}
	userService := service.UserService{Store: ust, LiveSt: lgst}
	boggleService := service.BoggleService{Store: gst, Bst: bst}
	validation.Setup(boggleService, userService)
	orch.Setup(boggleService, userService, judge)
	api.SetUp(apiMux)
}

func New() *Server {
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Boggle Api",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"local": "localhost:8080", "heroku": "https://protected-ravine-41774.herokuapp.com"},
	})
	router := mux.NewRouter()
	//router.Use(middleware.AddAuthHeaderMiddleware)
	host := viper.GetString("SERVER_HOST")
	port := viper.GetString("SERVER_PORT")
	addr := host + ":" + port
	s := Server{router, addr}
	s.SetupComponents()
	return &s
}

func (s Server) ServeHTTP() {
	srv := &http.Server{
		Handler:      middleware.Handle(s.Router),
		Addr:         s.Address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	logrus.Info("Server starting at addr: ", s.Address)
	logrus.Fatal(srv.ListenAndServe())
}
