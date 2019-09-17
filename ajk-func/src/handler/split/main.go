package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/theboss/ajk-emoji/ajk-func/src/infrastructure"
	"github.com/theboss/ajk-emoji/ajk-func/src/usecase"
)

var (
	store        = infrastructure.NewStorage()
	splitUsecase = usecase.NewSplit(store)
)

func handler(ctx context.Context, req events.S3Event) error {
	log.Printf("!!!SPLIT!!! req=%#v", req)
	// var reqBody model.RequestBody
	// if err := json.Unmarshal([]byte(req.Body), &reqBody); err != nil {
	// 	log.Print(err)
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: http.StatusBadRequest,
	// 		Headers:    constant.ResponseHeaders,
	// 		Body:       err.Error(),
	// 	}, nil
	// }

	// err := splitUsecase.SplitAndPut(
	// 	req.RequestContext.RequestID,
	// 	&reqBody,
	// )
	// if err != nil {
	// 	log.Print(err)
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// b, err := json.Marshal(struct {
	// 	ReqID string `json:"requestId"`
	// }{
	// 	ReqID: req.RequestContext.RequestID,
	// })
	// if err != nil {
	// 	log.Print(err)
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	return nil
}

func main() {
	lambda.Start(handler)
}
