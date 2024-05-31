package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := user{
			"joao",
			"joao.pedro@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}