package handler

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"syncmemo/entity"
	"syncmemo/model"
	"syncmemo/repository/request"
	"syncmemo/sercice"
	"syncmemo/store"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Casual struct {
	DB        *mongo.Database
	Validator *validator.Validate
	Kvs       *redis.Client
}

func (c *Casual) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var regInfo request.CasualUser
	if err := json.Unmarshal(reqBody, &regInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := c.Validator.Struct(regInfo); err != nil {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	err := model.EmailCore(ctx, c.DB, regInfo.Mail)
	if err != nil {
		if err == mongo.ErrNoDocuments { //メールアドレスが存在しない
			token := randNum(9999, 1000)
			pas, err := bcrypt.GenerateFromPassword([]byte(regInfo.Password), bcrypt.DefaultCost)
			if err != nil {
				RespondJSON(ctx, rw, &ErrResponse{
					Message: err.Error(),
				}, http.StatusInternalServerError)
				return
			}
			data := entity.User{
				Name:     regInfo.Mail,
				Mail:     regInfo.Mail,
				Password: string(pas),
			}
			if err := store.TokenSet(ctx, c.Kvs, data, token); err != nil {
				RespondJSON(ctx, rw, &ErrResponse{
					Message: err.Error(),
				}, http.StatusInternalServerError)
				return
			}

			if err := sercice.Send(regInfo.Mail, token); err != nil {
				RespondJSON(ctx, rw, &ErrResponse{
					Message: err.Error(),
				}, http.StatusInternalServerError)
				return
			}
			RespondJSON(ctx, rw, nil, http.StatusOK)
		} else { // 存在しない以外のエラー
			RespondJSON(ctx, rw, &ErrResponse{
				Message: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
	} else { //すでにメールアドレスが存在している
		RespondJSON(ctx, rw, nil, http.StatusBadRequest)
		return
	}

}
func randNum(max int, min int) string {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return strconv.Itoa(result)
}
