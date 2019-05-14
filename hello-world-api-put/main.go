package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type (
	helloWorld struct {
		Message string `json:"message"`
	}
)

// Handler lambda main function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var h helloWorld

	_ = json.Unmarshal([]byte(request.Body), &h)

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + h.Message,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)

}
