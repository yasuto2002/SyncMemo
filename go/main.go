package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	db, client := store.Connect(cfg.Mongo, ctx)
	defer func() {
		cancel()
		client.Disconnect(ctx)
	}()
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	chatroom := r.PathPrefix("/chatroom").Subrouter()
	chatroom.HandleFunc("/create/{name}", CR(CreateChatroom)).Methods(http.MethodPost)
	chatroom.HandleFunc("/list", CR(ListChatroom)).Methods(http.MethodGet)
	chatroom.HandleFunc("/connect", clientMW(chatroomWSHandler))

	cl := r.PathPrefix("/client").Subrouter()
	cl.HandleFunc("/list", ListAllClients).Methods(http.MethodGet)

	t := &handler.Test{DB: db, CTX: ctx}
	r.HandleFunc("/test", t.ServeHTTP).Methods(http.MethodGet)

	b := &handler.MakeBoard{DB: db}
	r.HandleFunc("/makeBoard", b.ServeHTTP).Methods(http.MethodPost)

	bl := &handler.BoardList{DB: db}
	r.HandleFunc("/boardList", bl.ServeHTTP).Methods(http.MethodPost)

	log.Println("Registered Handlers")

	log.Printf("Started Server on port : %v", port)
	headers := handlers.AllowedHeaders([]string{"*", "Content-Type", "*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methodsOk)(r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
