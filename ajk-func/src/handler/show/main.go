package main

import (
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/theboss/ajk-emoji/ajk-func/src/constant"
	"github.com/theboss/ajk-emoji/ajk-func/src/infrastructure"
	"github.com/theboss/ajk-emoji/ajk-func/src/usecase"
)

var (
	store       = infrastructure.NewStorage()
	showUsecase = usecase.NewShow(store)
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rc, err := showUsecase.GetEmojis(req.PathParameters["requestId"])
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{}, err
	}

	defer rc.Close()
	b, err := ioutil.ReadAll(rc)
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
