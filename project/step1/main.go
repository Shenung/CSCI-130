package main

import (
	"html/template"
	"log"
	"net/http"
)

func serve(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
