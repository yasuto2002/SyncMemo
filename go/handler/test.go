package handler

import (
	"context"
	"net/http"
	"syncmemo/auth"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	DB        *mongo.Database
	CTX       context.Context
	Validator *validator.Validate
	JWT       *auth.JWTer
}

type TestHandler interface {
	ServeHTTP()
}

type TestData struct {
	A string
}

func (t *Test) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	RespondJSON(ctx, rw, nil, http.StatusOK)
}
