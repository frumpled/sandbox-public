package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
)

type LambdaHandler = func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func createHandler(serveHTTP http.HandlerFunc) LambdaHandler {
	handler := func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		reqAccessorV2 := core.RequestAccessor{}
		req, err := reqAccessorV2.ProxyEventToHTTPRequest(event)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
			}, nil
		}

		log.Println(req.Method, req.URL.String())
		writer := core.NewProxyResponseWriter()
		req.URL.Path = fmt.Sprintf("/%s", event.PathParameters["proxy"])

		serveHTTP(writer, req.WithContext((ctx)))
		res, err := writer.GetProxyResponse()
		return res, err
	}

	return handler
}
