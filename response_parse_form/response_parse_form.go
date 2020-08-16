package main

import (
	"fmt"
	"net/http"
)

// get a form data when post form
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
