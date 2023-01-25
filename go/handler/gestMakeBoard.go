package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syncmemo/clock"
	"syncmemo/model"
	"syncmemo/repository/request"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type GestMakeBoard struct {
	DB        *mongo.Database
	Validator *validator.Validate
	Clock     clock.Clocker
}

func (G *GestMakeBoard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item request.Make
	if err := json.Unmarshal(reqBody, &item); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := G.Validator.Struct(item); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	id, err := model.MakeBords(ctx, G.DB, item, G.Clock)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, id, http.StatusOK)
}