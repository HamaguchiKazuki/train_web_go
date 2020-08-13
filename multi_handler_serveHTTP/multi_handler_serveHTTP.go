package main

import (
	"fmt"
	"net/http"
)

// HelloHandler is simple return "Hello!"
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// WorldHandler is simple return "World!"
type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
