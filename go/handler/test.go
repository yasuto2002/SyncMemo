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
	// ses()
	// model.MakePodcast(t.CTX, t.DB)
	u := entity.User{Mail: id}
	// b, err := t.JWT.GenerateToken(ctx, u)
	// if err != nil {
	// 	RespondJSON(ctx, rw, &ErrResponse{
	// 		Message: err.Error(),
	// 	}, http.StatusInternalServerError)
	// 	return
	// }
	// test := TestData{A: string(b)}
	//fmt.Errorf("user_id not found")
	RespondJSON(ctx, rw, u, http.StatusOK)
}
