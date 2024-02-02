package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Aff    string
	Alamat string
}

type People struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetDetailInfo() string {
	return "Memiliki 31 Orang"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var people = People{
			Name:    "Sholahuddin",
			Gender:  "Male",
			Hobbies: []string{"Membaca Buku", "Traveling", "Testing"},
			Info:    Info{"Testing Info 1", "Testing Info 2"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, people); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server Start at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
