package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Greeting string `json:"greeting"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	wsConn, _ := wsUpgrader.Upgrade(w, r, nil)
	defer wsConn.Close()
	//
	for {
		var msg Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("err JSON%v", err.Error())
		}
		fmt.Println(msg.Greeting)
	}
}

func main() {
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
