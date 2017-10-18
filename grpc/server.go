package grpc

import (
	"fmt"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

// Server ..
type Server interface {
	Listen(addr string) error
}

type server struct {
	logger pkg.Logger
}

// NewServer ..
func NewServer(logger pkg.Logger) Server {
	return &server{
		logger: logger,
	}
}

func (s *server) Listen(addr string) error {
	s.logger.Log(
		"msg", "grpc.server.listen",
		"addr", addr,
	)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	fooHandler := newFooServer()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(ui),
		grpc.StreamInterceptor(si),
	)
	RegisterFooServer(server, fooHandler)
	return server.Serve(listener)
}

func ui(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func(begin time.Time) {
		fmt.Println("UnaryInterceptor", time.Since(begin))
	}(time.Now())

	return handler(ctx, req)
}

func si(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func(begin time.Time) {
		fmt.Println("StreamInterceptor", time.Since(begin))
	}(time.Now())

	return handler(srv, ss)
}
