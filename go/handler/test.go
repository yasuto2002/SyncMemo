package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"syncmemo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	DB  *mongo.Database
	CTX context.Context
}

type TestHandler interface {
	ServeHTTP()
}

type TestData struct {
	A string
	B int
}

func (t *Test) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	model.MakePodcast(t.CTX, t.DB)
	test := TestData{A: "a", B: 2}
	data, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}
