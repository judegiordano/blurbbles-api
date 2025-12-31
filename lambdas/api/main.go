package main

import (
	"context"

	"blurbbles/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	server := internal.Server()
	app := adaptor.New(server)
	return app.ProxyWithContextV2(ctx, req)
}

func main() {
	lambda.Start(handler)
}
