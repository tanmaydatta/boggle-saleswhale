package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tanmaydatta/boggle/internal/api"
	"net/http"
	"time"
	"github.com/betacraft/yaag/yaag"
	"github.com/betacraft/yaag/middleware"

)

type Server struct {
	*mux.Router
	Address string
}

func (s Server) SetupComponents() {
	apiMux := s.PathPrefix("/api").Subrouter()
	api.SetUp(apiMux)
}

func New() *Server {
	yaag.Init(&yaag.Config{
		On: true,
		DocTitle: "Boggle Api",
		DocPath: "apidoc.html",
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

