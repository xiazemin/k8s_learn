package main

import (
	"fmt"
	"net/http"
)

func router(name string) {
	http.HandleFunc(name, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, name)
	})
}
func main() {
	router("/apple")
	router("/apple/banana")
	router("/apple1/banana")
	router("/apple1/banana/abc")
	router("/apple1/banana/abc/def")
	fmt.Println(http.ListenAndServe(":5678", nil))
}

//% go run main.go
//% curl http://127.0.0.1:5678/apple
//apple
