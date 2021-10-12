package server

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/frumpled/sandbox/twirp/proto/generated"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {

	resp := &pb.HelloResp{Text: "Hello " + req.Subject}

	fmt.Printf("%+v -> %+v\n", req, resp)

	return resp, nil
}

// Run the implementation in a local server
func main() {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8888", mux)
}
