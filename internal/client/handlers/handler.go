package handlers

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func NewHandler() Handler {
	return Handler{}
}

type Handler struct {
}

func (h Handler) Request(metric string, name string, val string) {
	var b strings.Builder
	b.WriteString("http://localhost:8080/update/")
	b.WriteString(metric)
	b.WriteString("/")
	b.WriteString(name)
	b.WriteString("/")
	b.WriteString(val)
	uri := b.String()

	client := http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if err != nil {
		panic(err)
	}
}
