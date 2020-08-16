package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	// 	w.Header().Set("Set-Cookie", c1.String())
	// 	w.Header().Add("Set-Cookie", c2.String())
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// h := r.Header.Get("Cookie")
	// h := r.Header["Cookie"]
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cookies := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cookies)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
