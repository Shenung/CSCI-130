package main

import "net/http"

func handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Can You See Me.\n"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
