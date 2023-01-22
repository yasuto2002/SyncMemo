package handler

import (
	"fmt"
	"net/http"
	"syncmemo/auth"
	"syncmemo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type BoardList struct {
	DB *mongo.Database
}

func (B *BoardList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}
	d, err := model.GetBoardList(ctx, B.DB, id)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, d, http.StatusOK)
}
