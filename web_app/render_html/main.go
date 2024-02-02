package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var isi = map[string]interface{}{
			"title": "Testing Golang Web",
			"name":  "Sholahuddin",
		}

		err = tmpl.Execute(w, isi)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}
