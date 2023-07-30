package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	log.Println("Fire it up!")

	index := template.Must(template.ParseFiles("templates/index.gohtml"))

	data := TodoPageData{
		PageTitle: "TODO Stuff",
		Todos: []Todo{
			{Title: "Ran code without it breaking initially", Done: true},
			{Title: "Clicked this and marked it as done", Done: false},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := index.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
