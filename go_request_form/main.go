package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseurl = "http://localhost:8080"

type mhs struct {
	ID   string
	Nama string
	Nim  int
}

func ambilidmhs(ID string) (mhs, error) {
	var err error
	var client = &http.Client{}
	var data mhs

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseurl+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var mhs1, err = ambilidmhs("D003")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("id: %s\t Nama: %s\t NIM: %d\n", mhs1.ID, mhs1.Nama, mhs1.Nim)
}
