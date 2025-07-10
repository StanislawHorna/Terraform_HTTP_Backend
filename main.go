package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"terraform_http_backend/src/handler"
	"terraform_http_backend/src/log"
)

func main() {
	log.Init()
	r := chi.NewRouter()
	r.Get("/state/{projectName}", handler.GetState)
	r.Post("/state/{projectName}", handler.SetState)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
