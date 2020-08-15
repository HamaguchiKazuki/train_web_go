package main

import (
	"fmt"
	"net/http"
)

// get a body when post
func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body) // body := [request body]

	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
