package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/frumpled/sandbox/twirp/proto/generated"
)

func main() {
	client := pb.NewHelloWorldProtobufClient(
		// "http://localhost:8888",
		"https://xxxxxxxxxx.execute-api.us-east-1.amazonaws.com/serverless_lambda_stage",
		&http.Client{},
	)

	request := &pb.HelloReq{
		Subject: "World",
		Bigness: 10,
	}
	fmt.Printf("Request:\n%+v\n", request)

	resp, err := client.Hello(context.Background(), request)
	if err == nil {
		fmt.Println(resp.Text) // prints "Hello World"
	} else {
		fmt.Printf("Error: %+v\n", err)
	}
}
