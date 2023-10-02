package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewSender() Handler {
	return Handler{client: http.Client{Timeout: 3 * time.Second}}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandleRequest
type HandleRequest interface {
	Request(metric string, name string, val string, port string)
}

type Handler struct {
	client http.Client
}

func (h *Handler) Request(metric, name, val, port string) {
	uri := fmt.Sprintf("http://%s/update/%s/%s/%s", port, metric, name, val)
	log.Println("uri", uri)

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

}
