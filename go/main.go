package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"syncmemo/config"
	"syncmemo/handler"
	"syncmemo/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	cfg, err := config.New()
	if err != nil {
		fmt.Println(err)
	}

	log.SetOutput(os.Stdout)

	flag.StringVar(&port, "p", cfg.Port, "port")
	flag.Parse()

	ctx, db, close := store.Connect(cfg.Mongo)
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
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(r))
}
