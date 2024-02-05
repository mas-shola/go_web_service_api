package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pahlawan struct {
	Nama  string
	Alias string
	Teman []string
}

func (p Pahlawan) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var people = Pahlawan{
			Nama:  "Sholahuddin",
			Alias: "Developer",
			Teman: []string{"Lantai 2", "Lantai 3", "Lantai 4"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, people); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server Started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
