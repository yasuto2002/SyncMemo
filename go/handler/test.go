package handler

import (
	"context"
	"fmt"
	"net/http"
	"syncmemo/auth"
	"syncmemo/entity"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	DB        *mongo.Database
	CTX       context.Context
	Validator *validator.Validate
	JWT       *auth.JWTer
}

type TestHandler interface {
	ServeHTTP()
}

type TestData struct {
	A string
}

const region = "ap-northeast-1"

var (
	fromEmailAddress = "test@yasuto0101.com"
	toEmailAddress   = "fujiya0101@gmail.com"
	subject          = "hey"
	body             = "hello world"
)

func (t *Test) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}
	// ses()
	// model.MakePodcast(t.CTX, t.DB)
	u := entity.User{Mail: id}
	// b, err := t.JWT.GenerateToken(ctx, u)
	// if err != nil {
	// 	RespondJSON(ctx, rw, &ErrResponse{
	// 		Message: err.Error(),
	// 	}, http.StatusInternalServerError)
	// 	return
	// }
	// test := TestData{A: string(b)}
	//fmt.Errorf("user_id not found")
	RespondJSON(ctx, rw, u, http.StatusOK)
}

func ses() {
	ctx := context.Background()

	// sdk API Client 作成
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		// error handling
	}
	client := sesv2.NewFromConfig(cfg)

	// SES API に投げ込むパラメタを作る
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &fromEmailAddress,
		Destination: &types.Destination{
			ToAddresses: []string{toEmailAddress}, // 配列なので複数指定可能
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &body, // 本文
					},
				},
				Subject: &types.Content{
					Data: &subject, // 件名
				},
			},
		},
	}

	// メール送信
	res, err := client.SendEmail(ctx, input)
	if err != nil {
		// error handling
	}
	fmt.Println(res.MessageId)
	fmt.Println("success!")
}
