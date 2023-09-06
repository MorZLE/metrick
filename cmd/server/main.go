package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	UpServer()

}

type Metric struct {
	metric string
	name   string
	value  string
}

func updatePage(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	metric := vars["metric"]
	if err != nill {
		res.WriteHeader(http.StatusNotFound)
	}
	if metric != "gauge" || "counter" {
		res.WriteHeader(http.StatusBadRequest)
	}
	name := vars["name"]
	value := vars["value"]
	m := Metric{metric, name, value}
	fmt.Println(metric, name, value)
}

func UpServer() {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{metric}/{name}/{value}`, updatePage)
	http.Handle("/", router)

	err := http.ListenAndServe(`:8080`, nil)
	if err == nil {
		panic(fmt.Errorf("ошибка запуска сервера: %d", err))
	}

}
