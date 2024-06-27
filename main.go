// starting point of Go program
package main

//packages imported to execute
import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

// base URL for Pokemon API
const apiURL = "https://pokeapi.co/api/v2/pokemon/"

// Defines a strut to map the JSON response from the API
type Pokemon struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

// Function to fetch a Pokemon by name
// *Pokemon is directing towards struct
func getPokemon(name string) (*Pokemon, error) {
	resp, err := http.Get(apiURL + name)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func main() {
	pokemonName:="mew"
	pokemon, err := getPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}

		fmt.Printf("Name: %s, ID: %d\n", pokemon.Name, pokemon.ID)
}
