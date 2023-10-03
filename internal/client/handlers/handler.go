package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MorZLE/metrick/internal/client/constants"
	"log"
	"net/http"
	"time"
)

func NewSender() Handler {
	return Handler{client: http.Client{Timeout: 3 * time.Second}}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandleRequest
type HandleRequest interface {
	Request(obj constants.Metrics, port string)
}

type Handler struct {
	client http.Client
}

func (h *Handler) Request(obj constants.Metrics, port string) {
	uri := fmt.Sprintf("http://%s/update", port)
	log.Println("uri", uri)
	body, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

}
