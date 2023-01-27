package handler

import (
	"fmt"
	"net/http"
	"syncmemo/auth"
	"syncmemo/store"

	"github.com/go-redis/redis/v9"
)

type Logout struct {
	JWTer *auth.JWTer
	Kvs   *redis.Client
}

func (logout *Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := logout.JWTer.GetToken(ctx, logout.Kvs, r)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := store.LoginDelete(ctx, logout.Kvs, token.JwtID()); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("not token in redis").Error(),
		}, http.StatusUnauthorized)
		return
	}
	RespondJSON(ctx, rw, nil, http.StatusOK)
}
