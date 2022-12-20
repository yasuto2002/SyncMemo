package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syncmemo/model"
	"syncmemo/repository/request"
	"syncmemo/repository/response"

	"go.mongodb.org/mongo-driver/mongo"
)

type MakeBoard struct {
	DB  *mongo.Database
	CTX context.Context
}

func (B *MakeBoard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item request.Make
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	id := model.MakeBords(B.CTX, B.DB, item)
	re := response.Make{ID: id}
	data, err := json.Marshal(re)
	if err != nil {
		fmt.Println(err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}
