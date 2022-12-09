package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var dataCh chan Data

type Data struct {
	MyData string `json:"myData"`
}

func main() {
	dataCh = make(chan Data, 1)

	http.Handle("/ws", websocket.Handler(dataHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dataHandler(ws *websocket.Conn) {
	for data := range dataCh {
		err := websocket.JSON.Send(ws, data)
		if err != nil {
			log.Printf("error sending data: %v\n", err)
			return
		}
	}
}
