package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type ApiServer struct {}

//START
func (s *ApiServer) Start() error {
	var ro = mux.NewRouter()
	ro.HandleFunc("/", s.Hello())
	
	
	return http.ListenAndServe(":3000", ro)
}

func (s *ApiServer) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "you")
	}
}

//main
func main() {
	var s *ApiServer
	if err := s.Start(); err != nil {
		log.Fatal("start",err)
	}
	log.Println(viper.GetString("port"))
	//init yml
	if err := initViper(); err != nil {
		log.Fatal(err)
	}
}

func initViper() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
