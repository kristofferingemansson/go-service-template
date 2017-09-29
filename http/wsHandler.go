package http

import (
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/kristofferingemansson/go-service-template/pkg"
	"github.com/kristofferingemansson/go-service-template/quote"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type wsHandler struct {
	logger  pkg.Logger
	service quote.Service
}

// NewWsHandler ..
func NewWsHandler(logger pkg.Logger, service quote.Service) RoutableHandler {
	return &wsHandler{
		logger:  logger,
		service: service,
	}
}

func (h *wsHandler) Route(router chi.Router) {
	router.Get("/", h.upgradeConnection)
}

func (h *wsHandler) upgradeConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		encodeError(w, err)
		return
	}

	go h.handleConnection(conn)
}

func (h *wsHandler) handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	h.logger.Log(
		"msg", "http.wsHandler.handleConnection: Connected",
	)

	defer h.logger.Log(
		"msg", "http.wsHandler.handleConnection: Disconnected",
	)

	messages := make(chan []byte)
	go func() {
		for {
			if t, r, err := conn.NextReader(); err != nil {
				close(messages)
				return
			} else if t == websocket.TextMessage {
				if b, err := ioutil.ReadAll(r); err == nil { // NB!
					messages <- b
				}
			}
		}
	}()

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	var writeLock sync.Mutex

	for {
		select {
		case <-t.C:
			if q, err := h.service.GenerateQuote(); err != nil {
				return
			} else {
				writeLock.Lock()
				conn.WriteJSON(map[string]interface{}{
					"quote": q,
				})
				writeLock.Unlock()
			}
		case _, ok := <-messages:
			if !ok {
				return
			}
		}
	}

}
