package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type country struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type allCountries []country

var countrys = allCountries{
	{
		ID:          "1",
		Name:        "Bangladesh",
		Description: "Bangladesh is a very beautiful country surrounding with river and natural beauty",
	},
	{
		ID:          "2",
		Name:        "India",
		Description: "India is a very dynamic economic and multicultural country. I have a plan to visit India",
	},
	{
		ID:          "2",
		Name:        "USA",
		Description: "Most young people dream to study in USA. It's beautiful country",
	},
}

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
