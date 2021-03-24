package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/apple", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "apple")
	})
	fmt.Println(http.ListenAndServe(":5678", nil))
}

//% go run main.go
//% curl http://127.0.0.1:5678/apple
//apple
