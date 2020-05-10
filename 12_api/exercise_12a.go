package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//BaseURL of the api
const BaseURL = "https://swapi.dev/api"

// Planet contains planet details
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person contains details of a person
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople contains list of person
type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Println("Error fetching HomeWorld:", err)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)

}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "/people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	fmt.Println(w, res)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("faled to parse request body")
	}

	var people AllPeople
	fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing error", err)
	}
	fmt.Println(people)

	for _, pers := range people.People {
		pers.getHomeWorld()
		fmt.Println(pers)
	}
}

func main() {
	fmt.Print(BaseURL)
	http.HandleFunc("/people", getPeople)
	fmt.Print("Server started on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
