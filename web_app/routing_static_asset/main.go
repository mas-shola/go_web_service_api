package main

import (
	"fmt"
	"net/http"
	"time"
)

func logRequestStatis(r *http.Request) {
	fmt.Printf("[%s] %s %s (/statis)\n", time.Now().Format("2001-01-01 15:01:01"), r.Method, r.URL.Path)
}

func main() {
	// http.Handle("/statis/", http.StripPrefix("/statis/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/statis", func(w http.ResponseWriter, r *http.Request) {
		logRequestStatis(r)
		http.StripPrefix("/statis", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
	})
	// logRequestStatis(r)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
