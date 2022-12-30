package handler

import (
	"net/http"
	"syncmemo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type BoardList struct {
	DB *mongo.Database
}

func (B *BoardList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	d, err := model.GetBoardList(ctx, B.DB)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, d, http.StatusOK)
}
