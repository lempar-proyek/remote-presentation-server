package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This application is work")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Upgrading to websocket...")
}

func setupRoutes() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Starting Application...")
	setupRoutes()
	fmt.Println("Server will run on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
