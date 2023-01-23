package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syncmemo/auth"
	"syncmemo/clock"
	"syncmemo/model"
	"syncmemo/repository/request"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type MakeBoard struct {
	DB        *mongo.Database
	Validator *validator.Validate
	Clock     clock.Clocker
}

func (B *MakeBoard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mail, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item request.Make
	if err := json.Unmarshal(reqBody, &item); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := B.Validator.Struct(item); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	item.Mail = mail

	id, err := model.MakeBords(ctx, B.DB, item, B.Clock)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, id, http.StatusOK)
}
