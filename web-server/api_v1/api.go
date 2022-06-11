package api_v1

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Fprintf(w, "Path Not Found %d", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		fmt.Fprintf(w, "URL Not Found %d", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, Asad — From RYK")
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		fmt.Fprintf(w, "Path Not Found %d", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		fmt.Fprintf(w, "URL Not Found %d", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %v \n", name)
	fmt.Fprintf(w, "Address: %v \n", address)
}
