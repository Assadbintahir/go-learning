package main

import (
	"net/http"
	"web-server/api_v1"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", Logger(api_v1.HandleHello))
	http.HandleFunc("/form", Logger(api_v1.HandleForm))

	logger.Printf("Server starting at port %v \n", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		logger.Fatalf("Internal Server Error: %v", err)
	}
}
