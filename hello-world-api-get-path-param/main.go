package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler lambda main function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello World " + request.PathParameters["id"],
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)

}
