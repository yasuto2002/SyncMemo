package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"syncmemo/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type BoardList struct {
	DB  *mongo.Database
	CTX context.Context
}

func (B *BoardList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	d := model.GetBoardList(B.CTX, B.DB)
	// re := response.Make{ID: id}
	// j := response.BoardList{Boards: d}
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(rw, "%s", data); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
