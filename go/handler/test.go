package handler

import (
	"context"
	"fmt"
	"net/http"
	"syncmemo/auth"
	"syncmemo/entity"

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
	id, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}
	u := entity.User{Mail: id}
	RespondJSON(ctx, rw, u, http.StatusOK)
}
