package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/header", headers)
	server.ListenAndServe()
}
