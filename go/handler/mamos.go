package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syncmemo/model"
	"syncmemo/repository/request"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Memos struct {
	Validator *validator.Validate
	DB        *mongo.Database
}

func (memos *Memos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var memosInfo request.Memos
	if err := json.Unmarshal(reqBody, &memosInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := memos.Validator.Struct(memosInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	resMemo, err := model.GetMemos(ctx, memos.DB, memosInfo.Id)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	RespondJSON(ctx, rw, resMemo, http.StatusOK)
}
