package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syncmemo/auth"
	"syncmemo/model"
	"syncmemo/repository/request"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardDelete struct {
	DB        *mongo.Database
	Validator *validator.Validate
}

func (BD *BoardDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mail, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteInfo request.BoardDelete
	if err := json.Unmarshal(reqBody, &deleteInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := BD.Validator.Struct(deleteInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if err := model.DeleteBoards(ctx, BD.DB, deleteInfo.BoardId, mail); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	RespondJSON(ctx, rw, nil, http.StatusOK)
}
