package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseurl = "http://localhost:8080"

type mhs struct {
	ID   string
	Nama string
	Nim  int
}

func ambilmhs() ([]mhs, error) {
	var err error
	var client = &http.Client{}
	var data []mhs

	request, err := http.NewRequest("GET", baseurl+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var users, err = ambilmhs()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("id: %s\t Nama: %s\t NIM: %d\n", each.ID, each.Nama, each.Nim)
	}
}
