package main

import (
	"io"
	"net/http"
	"strconv"
)

func cV(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	cookie, err := r.Cookie("cookie-value")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cookie-value",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	//io.WriteString("count:")

	io.WriteString(w, "value:"+cookie.Value)
}
func main() {
	http.HandleFunc("/", cV)
	http.ListenAndServe(":8080", nil)
}
