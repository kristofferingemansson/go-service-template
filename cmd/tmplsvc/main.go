package main

import (
	"flag"
	"github.com/kristofferingemansson/go-service-template/grpc"
	"github.com/kristofferingemansson/go-service-template/http"
	"github.com/kristofferingemansson/go-service-template/inmem"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"github.com/kristofferingemansson/go-service-template/quote"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8081", "http listen address")
		grpcAddr = flag.String("grpc.addr", ":8082", "grpc listen address")
		webRoot  = flag.String("webroot", "./www", "path to static files root")
	)

	flag.Parse()

	var errors = make(chan error)
	logger := pkg.StdLogger

	quoteRepository := inmem.NewQuoteRepository()
	quoteService := quote.NewService(quoteRepository)

	go func() {
		staticHandler := http.NewStaticHandler(logger, *webRoot)
		apiHandler := http.NewAPIHandler(logger, quoteService)
		wsHandler := http.NewWsHandler(logger, quoteService)

		server := http.NewServer(logger)
		errors <- server.Listen(*httpAddr, staticHandler, apiHandler, wsHandler)
	}()

	go func() {
		server := grpc.NewServer(logger)
		errors <- server.Listen(*grpcAddr)
	}()

	if err, ok := <-errors; ok {
		logger.Log(
			"msg", "error",
			"error", err.Error(),
		)
	}
}
