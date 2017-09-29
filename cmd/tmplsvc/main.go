package main

import (
	"flag"
	"github.com/kristofferingemansson/go-service-template/http"
	"github.com/kristofferingemansson/go-service-template/pkg"
)

func main() {
	var (
		addr    = flag.String("addr", ":8080", "http listen address")
		webroot = flag.String("webroot", "./www", "path to static files root")
	)

	flag.Parse()

	var errors = make(chan error)
	logger := pkg.StdLogger

	go func() {
		staticHandler := http.NewStaticHandler(logger, *webroot)
		apiHandler := http.NewAPIHandler(logger)
		wsHandler := http.NewWsHandler(logger)

		server := http.NewServer(logger)
		errors <- server.Listen(*addr, staticHandler, apiHandler, wsHandler)
	}()

	if err, ok := <-errors; ok {
		logger.Log(
			"msg", "error",
			"error", err.Error(),
		)
	}
}
