package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (Response, error) {
	var r Request
	if err := json.Unmarshal([]byte(req.Body), &r); err != nil {
		return Response{Message: "Invalid input"}, nil
	}

	msg := fmt.Sprintf("Hello, %s!", r.Name)
	return Response{Message: msg}, nil
}

func main() {
	lambda.Start(handler)
}
