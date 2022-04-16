package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type country struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type allCountries []country

var countries = allCountries{
	{
		ID:          "1",
		Name:        "Bangladesh",
		Description: "Bangladesh is a very beautiful country surrounding with river and natural beauty",
	},
	// {
	// 	ID:          "2",
	// 	Name:        "India",
	// 	Description: "India is a very dynamic economic and multicultural country. I have a plan to visit India",
	// },
	// {
	// 	ID:          "2",
	// 	Name:        "USA",
	// 	Description: "Most young people dream to study in USA. It's beautiful country",
	// },
}

// Create a New Country
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newCountry country
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Create a new country .Please Insert Country Name and Description.")
	}

	json.Unmarshal(reqBody, &newCountry)
	countries = append(countries, newCountry)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCountry)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I am From Home Routing!")
}

func main() {

	fmt.Println("Hello World")

	// Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)

	// Home Route , Url = /
	router.HandleFunc("/", HomePage)

	// Create a New Country Server
	router.HandleFunc("/create", createEvent).Methods("POST")

	// server setup and Running Server : 8080 port
	log.Fatal(http.ListenAndServe(":8080", router))

}
