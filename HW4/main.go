package main

import (
	//"io"
	"html/template"
	"log"
	"net/http"
)

func surf(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templ.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}
func main() {

	http.HandleFunc("/", surf)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8080", nil)
}
