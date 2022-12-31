package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"syncmemo/model"
	"syncmemo/repository/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type MakeBoard struct {
	DB *mongo.Database
}

func (B *MakeBoard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item request.Make
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	id, err := model.MakeBords(ctx, B.DB, item)
	if err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, id, http.StatusOK)
}
