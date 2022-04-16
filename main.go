package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I am FRom Home Routing!")
}

func main() {

	fmt.Println("Hello World")

	// Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)

	// Home Route , Url = /
	router.HandleFunc("/", HomePage)

	// server setup and Running Server : 8080 port
	log.Fatal(http.ListenAndServe(":8080", router))

}
