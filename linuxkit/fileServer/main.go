package main

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":9909", nil)
}
