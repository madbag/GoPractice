package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const APIURL = "http://www.omdbapi.com/?s=batman&apikey=YOUR_API_KEY"

type Movies struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(APIURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var data struct {
		Search []Movies `json:"Search"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Search)
}

func main() {
	http.HandleFunc("/movies", getMovies)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server started at http://localhost:8080")
}
