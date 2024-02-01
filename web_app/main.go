package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	var pesan = "Welcome"
	w.Write([]byte(pesan))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	var pesan = "HELLO WORLD"
	w.Write([]byte(pesan))
}

func logRequest(r *http.Request) {
	fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
}

func logRequestTesting(r *http.Request) {
	fmt.Printf("[%s] %s %s (from /testing)\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		logRequestTesting(r)
		w.Write([]byte("Hello Testing"))
	})
	var alamat = "localhost:9000"
	fmt.Printf("Server started at %s\n", alamat)
	err := http.ListenAndServe(alamat, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
