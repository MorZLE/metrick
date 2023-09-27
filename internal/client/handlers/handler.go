package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewHandler(addr string) Handler {
	return Handler{port: addr}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandleRequest
type HandleRequest interface {
	Request(metric string, name string, val string)
}

type Handler struct {
	port string
}

func (h *Handler) Request(metric, name, val string) {
	uri := fmt.Sprintf("http://%s/update/%s/%s/%s", h.port, metric, name, val)
	log.Println("uri", uri)
	client := http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}
