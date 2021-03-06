package main

import (
	"io"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	key := "n"
	val := req.FormValue(key)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<form method="GET">
		 <input type="text" name="n">
		 <input type="submit">
		</form>`+val)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
