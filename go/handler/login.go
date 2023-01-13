package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syncmemo/auth"
	"syncmemo/repository/request"
	"syncmemo/repository/response"
	"syncmemo/sercice"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Login struct {
	Validator *validator.Validate
	Kvs       *redis.Client
	DB        *mongo.Database
	JWTer     *auth.JWTer
}

func (log *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var logInfo request.Login
	if err := json.Unmarshal(reqBody, &logInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := log.Validator.Struct(logInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	user, err := sercice.Login(ctx, log.DB, logInfo.Mail, logInfo.Password)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	token, err := log.JWTer.GenerateToken(ctx, log.Kvs, user)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, response.Login{Token: string(token)}, http.StatusOK)
}
