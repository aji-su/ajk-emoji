package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/theboss/ajk-emoji/ajk-func/src/constant"
	"github.com/theboss/ajk-emoji/ajk-func/src/infrastructure"
	"github.com/theboss/ajk-emoji/ajk-func/src/usecase"
)

var (
	store           = infrastructure.NewStorage()
	downloadUsecase = usecase.NewDownload(store)
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url, err := downloadUsecase.Download(req.PathParameters["requestId"])
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{}, err
	}

	b, err := json.Marshal(struct {
		URL string `json:"url"`
	}{
		URL: url,
	})
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    constant.ResponseHeaders,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handler)
}
