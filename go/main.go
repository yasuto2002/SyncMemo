package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"syncmemo/auth"
	"syncmemo/clock"
	"syncmemo/config"
	"syncmemo/handler"
	"syncmemo/store"

	"github.com/go-playground/validator/v10"
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

	authKvs := store.NewAuthKvs(ctx, cfg.Kvs)
	loginKvs := store.NewLoginKvs(ctx, cfg.Kvs)

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	v := validator.New()

	clocker := clock.RealClocker{}
	jwter, err := auth.NewJWTer(clocker, loginKvs)
	chatroom := r.PathPrefix("/chatroom").Subrouter()
	chatroom.HandleFunc("/create/{name}", CR(CreateChatroom)).Methods(http.MethodPost)
	chatroom.HandleFunc("/list", CR(ListChatroom)).Methods(http.MethodGet)
	chatroom.HandleFunc("/connect", clientMW(chatroomWSHandler))

	cl := r.PathPrefix("/client").Subrouter()
	cl.HandleFunc("/list", ListAllClients).Methods(http.MethodGet)

	t := &handler.Test{DB: db, CTX: ctx, Validator: v, JWT: jwter}
	test := r.PathPrefix("/test").Subrouter()
	test.Use(handler.AuthMiddleware(jwter, loginKvs))
	test.HandleFunc("/", t.ServeHTTP).Methods(http.MethodGet)

	b := &handler.MakeBoard{DB: db, Validator: v}
	r.HandleFunc("/makeBoard", b.ServeHTTP).Methods(http.MethodPost)

	bl := &handler.BoardList{DB: db}
	r.HandleFunc("/boardList", bl.ServeHTTP).Methods(http.MethodPost)

	casual := &handler.Casual{DB: db, Validator: v, Kvs: authKvs}
	r.HandleFunc("/casual", casual.ServeHTTP).Methods(http.MethodPost)

	reg := &handler.Register{DB: db, Validator: v, Kvs: authKvs}
	r.HandleFunc("/register", reg.ServeHTTP).Methods(http.MethodPost)
	if err != nil {
		return
	}
	login := &handler.Login{DB: db, Validator: v, Kvs: loginKvs, JWTer: jwter}
	r.HandleFunc("/login", login.ServeHTTP).Methods(http.MethodPost)

	log.Println("Registered Handlers")
	log.Printf("Started Server on port : %v", port)
	headers := handlers.AllowedHeaders([]string{"*", "Content-Type", "*"})
	authorization := handlers.AllowedHeaders([]string{"*", "Authorization", "*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(headers, authorization, origins, methodsOk)(r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
