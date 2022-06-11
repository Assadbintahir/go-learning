package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
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

func handleForm(w http.ResponseWriter, r *http.Request) {
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

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", Logger(handleHello))
	http.HandleFunc("/form", Logger(handleForm))

	logger.Printf("Server starting at port %v \n", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		logger.Fatalf("Internal Server Error: %v", err)
	}
}
