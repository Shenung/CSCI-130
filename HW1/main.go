package main

import (
	"log"
	"os"
	"text/template"
)

type player struct {
	Name   string
	Health int
	Life   bool
}

func main() {
	player1 := player{
		Name:   "The Player Who Died A Painful Death",
		Health: 0,
		Life:   false,
	}

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, player1)
	if err != nil {
		log.Fatalln(err)
	}
}
