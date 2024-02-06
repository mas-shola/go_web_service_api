package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", ActionIndex)

	fmt.Println("Server Started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	content := [] struct {
		Nama string
		Umur int
	} {
		{"Testing 1", 24},
		{"Testing 2", 23,
		{"Testing 3", 21},
	}

	jsonInbytes, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(jsonInbytes)
}