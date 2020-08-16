package main

import (
	"fmt"
	"net/http"
)

// get a form data when post form
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, "1", r.FormValue("hello"))
	fmt.Fprintln(w, "2", r.PostFormValue("hello"))
	fmt.Fprintln(w, "3", r.PostForm)
	fmt.Fprintln(w, "4", r.MultipartForm)
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
