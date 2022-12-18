package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"syncmemo/handler"
	"syncmemo/model"
)

const (
	readBuffSize = 2 << 10
	writeBuffSize
)

const (
	nameHeader = "WS-NAME"
	idHeader   = "WS-ID"
)

var port string

func main() {

	log.SetOutput(os.Stdout)

	flag.StringVar(&port, "p", "8080", "port")
	flag.Parse()

	ctx, db, close := model.Connect()
	defer close()
	r := mux.NewRouter()

	chatroom := r.PathPrefix("/chatroom").Subrouter()
	chatroom.HandleFunc("/create/{name}", CR(CreateChatroom)).Methods(http.MethodPost)
	chatroom.HandleFunc("/list", CR(ListChatroom)).Methods(http.MethodGet)
	chatroom.HandleFunc("/connect", clientMW(chatroomWSHandler))

	client := r.PathPrefix("/client").Subrouter()
	client.HandleFunc("/list", ListAllClients).Methods(http.MethodGet)

	t := &handler.Test{DB: db, CTX: ctx}
	r.HandleFunc("/test", t.ServeHTTP).Methods(http.MethodGet)

	log.Println("Registered Handlers")

	log.Printf("Started Server on port : %v", port)
	http.ListenAndServe(":"+port, r)
}
