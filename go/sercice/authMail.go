package sercice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const region = "ap-northeast-1"

var (
	fromEmailAddress = "test@yasuto0101.com"
	subject          = "hey"
	body             = "hello world"
)

func Send(toEmailAddress string, body string) error {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return err
	}
	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &fromEmailAddress,
		Destination: &types.Destination{
			ToAddresses: []string{toEmailAddress},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &body,
					},
				},
				Subject: &types.Content{
					Data: &subject,
				},
			},
		},
	}

	res, err := client.SendEmail(ctx, input)
	if err != nil {
		return err
	}
	fmt.Println(res.MessageId)
	fmt.Println("success!")
	return nil
}
