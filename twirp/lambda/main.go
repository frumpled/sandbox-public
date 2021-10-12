package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

	pb "github.com/frumpled/sandbox/twirp/proto/generated"
	"github.com/frumpled/sandbox/twirp/server"
)

var lambdaHandler LambdaHandler

func init() {
	fmt.Printf("Cold start @ %+v", time.Now())

	twirpHandler := pb.NewHelloWorldServer(&server.HelloWorldServer{})
	lambdaHandler = createHandler(twirpHandler.ServeHTTP)

	fmt.Printf("Cold start finished @ %+v", time.Now())
}

func main() {
	lambda.Start(lambdaHandler)
}
