package apis

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	api    *mux.Router
	logger *logrus.Logger
}

func New() *ApiServer {
	return &ApiServer{
		api:    mux.NewRouter(),
		logger: logrus.New(),
	}
}
func init() {
	var s ApiServer
	s.logger.SetOutput(os.Stdout)
}

//START
func (s *ApiServer) Start() error {
	var r = mux.NewRouter()
	r.HandleFunc("/", s.Hello()).Methods("GET")

	if err := s.ConfigurateLoger(); err != nil {
		return err
	}
	s.logger.Info("start")
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:3000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *ApiServer) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "you")
	}
}

//loger

func (s *ApiServer) ConfigurateLoger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Println(err)
	}
	s.logger.SetLevel(level)
	return nil
}
