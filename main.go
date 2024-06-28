package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiURL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	Height uint   `json:"height"`
}

type ListOfPokemon struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getPokemon(name string) (*Pokemon, error) {
	//response & error
	resp, err := http.Get(apiURL + name)
	if err != nil {
		return nil, err
	}

	//Delay the execution of a statement until the end of a function
	defer resp.Body.Close()

	//if status is not OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	//decoding pokemon json file.
	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, fmt.Errorf("failed to decode pokemon data: %v", err)
	}

	return &pokemon, nil
}

// first 10 pokemon of the API
func getFirstTenPokemon() ([]*Pokemon, error) {
	resp, err := http.Get(apiURL + "?limit=10")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var listOfPokemon ListOfPokemon 
		if err:= json.NewDecoder(resp.Body).Decode(&listOfPokemon); err != nil {
			return nil, fmt.Errorf("failed to decode pokemon list: %v", err)
	}

	var pokemons []*Pokemon
	for _, p := range listOfPokemon.Results {
		pokemon, err := getPokemon(p.Name)
		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, pokemon)

	}
	return pokemons, nil
}


//last 10 pokemon 
func main() {
	pokemonName := "pikachu"
	pokemon, err := getPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result 1 = Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)

	pokemons, err := getFirstTenPokemon()
	if err != nil {
	    log.Fatalf("Error fetching the first 10 pokemon: %v", err)
	}

	for _, pokemon:= range pokemons {
		fmt.Printf("Result 2 = Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)
	}
}
