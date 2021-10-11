package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/frumpled/sandbox/twirp/proto/generated"
)

func main() {
	client := pb.NewHelloWorldProtobufClient("http://localhost:8888", &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
	if err == nil {
		fmt.Println(resp.Text) // prints "Hello World"
	} else {
		fmt.Printf("Error: %+v\n", err)
	}
}
