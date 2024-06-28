package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiURL = "https://pokeapi.co/api/v2/pokemon/"

// Defines a strut to map the JSON response from the API
type Pokemon struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	Height uint `json: "height"`
}

type PokemonListResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Function to fetch a Pokemon by name
func getPokemon(name string) (*Pokemon, error) {
	resp, err := http.Get(apiURL + name)
	if err != nil {
		return nil, err
	}
	// Ensure response body is closed after function completes
	defer resp.Body.Close()

	// Correct status check
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}

// List first 10 pokemon
func firstTenPokemon() ([]*Pokemon, error) {
	resp, err := http.Get(apiURL + "?limit=10")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var listResponse PokemonListResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResponse); err != nil {
		return nil, err
	}

	var pokemons []*Pokemon
	for _, result := range listResponse.Results {
	    pokemon, err:= getPokemon(result.Name)
	    if err != nil {
	        return nil, err
	    }
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func main() {
	pokemonName := "mew"
	pokemon, err := getPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Name: %s, ID: %d Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)

	pokemons, err := firstTenPokemon()
	if err != nil {
		log.Fatal(err)
	}

	for _, pokemon:= range pokemons {
		fmt.Printf("Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)
	}
}
