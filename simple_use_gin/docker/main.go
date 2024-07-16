package main

import (
	"net/http"
)

func sayer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello liwenzhou.com!"))
}

func main() {
	http.HandleFunc("/", sayer)
	http.ListenAndServe(":8080", nil)

}
