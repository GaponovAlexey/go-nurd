package main

import (
	"g/l/pkg/app/apis"
	"log"

	"github.com/spf13/viper"
)

//main
func main() {


	s := apis.New()
	if err := s.Start(); err != nil {
		log.Fatal("start", err)
	}



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
