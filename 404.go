package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", aa)
	http.HandleFunc("/car/", bb)

	http.ListenAndServe(":8080", nil)
}

func aa(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	io.WriteString(w, "Main path")
}

func bb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/car/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	io.WriteString(w, "Car path")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Page not found")
	}
}
