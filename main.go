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
func createCountry(w http.ResponseWriter, r *http.Request) {
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

// Return All Country
func getAllCountries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(countries)

}

// Return one country based on id
func getOneCountry(w http.ResponseWriter, r *http.Request) {
	countryID := mux.Vars(r)["id"]

	for _, singleCountry := range countries {
		if singleCountry.ID == countryID {
			json.NewEncoder(w).Encode(singleCountry)
		}
	}
}

// Update one country based on id
func updateCountry(w http.ResponseWriter, r *http.Request) {
	countryID := mux.Vars(r)["id"]
	var updatedcountry country

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the country title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedcountry)

	for i, singlecountry := range countries {
		if singlecountry.ID == countryID {
			singlecountry.Name = updatedcountry.Name
			singlecountry.Description = updatedcountry.Description
			countries = append(countries[:i], singlecountry)
			json.NewEncoder(w).Encode(singlecountry)
		}
	}
}

// Delete country based on country id
func deleteCountry(w http.ResponseWriter, r *http.Request) {
	countryID := mux.Vars(r)["id"]

	for i, singleCountry := range countries {
		if singleCountry.ID == countryID {
			countries = append(countries[:i], countries[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", countryID)
		}
	}
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
	router.HandleFunc("/create", createCountry).Methods("POST")

	// Return All Country
	router.HandleFunc("/countries", getAllCountries).Methods("GET")

	// Return One country based on country id
	router.HandleFunc("/countries/{id}", getOneCountry).Methods("GET")

	// Update One country based on country id
	router.HandleFunc("/countries/{id}", updateCountry).Methods("PATCH")

	// Delete country based on country id
	router.HandleFunc("/countries/{id}", deleteCountry).Methods("DELETE")

	// server setup and Running Server : 8080 port
	log.Fatal(http.ListenAndServe(":8080", router))

}
