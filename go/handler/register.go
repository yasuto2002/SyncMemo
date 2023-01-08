package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syncmemo/model"
	"syncmemo/repository/request"
	"syncmemo/store"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Register struct {
	Validator *validator.Validate
	Kvs       *redis.Client
	DB        *mongo.Database
}

func (reg *Register) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var regInfo request.Register

	if err := json.Unmarshal(reqBody, &regInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := reg.Validator.Struct(regInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	casual, err := store.TokenGet(ctx, reg.Kvs, regInfo.Mail)
	if err == redis.Nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	} else if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	if casual.Token != regInfo.Token {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := store.TokenDelete(ctx, reg.Kvs, regInfo.Mail); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err = model.Reg(ctx, reg.DB, &casual.User)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	RespondJSON(ctx, rw, nil, http.StatusOK)
}
