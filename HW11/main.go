package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		val := req.FormValue("name")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
			Enter your name:
		 <input type="text" name="name">
		 <input type="submit">
     </form>`+val)
	})
	http.ListenAndServe(":8080", nil)
}
