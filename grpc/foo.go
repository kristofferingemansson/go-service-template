package grpc

import (
	"golang.org/x/net/context"
)

type fooServer struct{}

func newFooServer() FooServer {
	return &fooServer{}
}

func (fooServer) GetFoo(ctx context.Context, request *FooRequest) (*FooResponse, error) {
	var ret FooResponse
	ret.Name = "Hello world!"
	return &ret, nil
}
