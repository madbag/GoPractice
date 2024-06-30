package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	apiURL          = "https://pokeapi.co/api/v2/pokemon/"
	totalPokemonURL = "https://pokeapi.co/api/v2/pokemon/"
)

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

type PokemonList struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// type PokemonType struct {
// 	Results []struct {
// 		Name  string `json:"name"`
// 		URL   string `json:"url"`
// 		Types []struct {
// 			Type string `json:"type"`
// 		} `json:"types"`
// 	}
// }

type TypeResponse struct {
	Pokemon []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon"`
}

// get pokemon based on the name
func getPokemon(name string) (*Pokemon, error) {
	//response & error
	resp, err := http.Get(apiURL + name)
	if err != nil {
		return nil, err
	}

	//Delay the execution of a statement until the end of the above function
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
	if err := json.NewDecoder(resp.Body).Decode(&listOfPokemon); err != nil {
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

// Get total Pokemon
func getTotalPokemonCount() (int, error) {
	resp, err := http.Get(totalPokemonURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Count int `json:"count"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode pokemon count: %v", err)
	}

	return result.Count, nil
}

// get last 10 pokemon
func getLastTenPokemon() ([]Pokemon, error) {
	count, err := getTotalPokemonCount()
	if err != nil {
		return nil, err
	}

	var pokemons []Pokemon
	limit := 10
	offset := count - limit
	url := fmt.Sprintf("%s?limit=%d&offset=%d", apiURL, limit, offset)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var listOfPokemon ListOfPokemon
	if err := json.NewDecoder(resp.Body).Decode(&listOfPokemon); err != nil {
		return nil, fmt.Errorf("failed to decode pokemon list:%v", err)
	}

	for _, p := range listOfPokemon.Results {
		pokemon, err := getPokemon(p.Name)
		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, *pokemon)
	}

	return pokemons, nil
}

// get random pokemon
func getRandomPokemon() (Pokemon, error) {
	rand.Seed(time.Now().UnixNano())

	randomID := rand.Intn(898) + 1

	url := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d", randomID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		log.Fatalf("Error decoding data: %v", err)
	}

	return pokemon, nil
	// fmt.Printf("Name: %s,\n ID:%d\n Height:%d\n", pokemon.Name, pokemon.ID, pokemon.Height)
}

// get 10 random pokemon
func getTenRandomPokemon() ([]Pokemon, error) {
	var pokemons []Pokemon

	for i := 0; i < 10; i++ {
		randomID := rand.Intn(898) + 1
		url := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d", randomID)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var pokemon Pokemon
		if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
			log.Fatalf("Error decoding data: %v", err)
		}
		pokemons = append(pokemons, pokemon)

	}

	return pokemons, nil
}

// Get pokemon acc to type
func getPokemonByType(pokemonType string) ([]Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/type/%s", pokemonType)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var typeResponse TypeResponse
	if err := json.NewDecoder(resp.Body).Decode(&typeResponse); err != nil {
		return nil, err
	}

	var pokemons []Pokemon
	for _, p := range typeResponse.Pokemon {
		pokemons = append(pokemons, p.Pokemon)
	}

	return pokemons, nil
}

func main() {
	http.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		pokemonName := "pikachu"
		pokemon, err := getPokemon(pokemonName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Result 1 = Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)

		fmt.Fprintf(w, "####################\n")

		pokemons, err := getFirstTenPokemon()
		if err != nil {
			log.Fatalf("Error fetching the first 10 pokemon: %v", err)
		}
		for _, pokemon := range pokemons {
			fmt.Fprintf(w, "Result 2 = Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)
		}

		fmt.Fprintf(w, "####################\n")

		morePokemons, err := getLastTenPokemon()
		if err != nil {
			log.Fatal(err)
		}
		for _, pokemon := range morePokemons {
			fmt.Fprintf(w, "Result 3 = Name:%s, ID:%d, Height:%d\n", pokemon.Name, pokemon.ID, pokemon.Height)
		}

		fmt.Fprintf(w, "####################\n")

		randomPokemon, err:= getRandomPokemon()
		if err != nil {
			log.Fatal(err) 
		}
		// for _, pokemon := range randomPokemon {
			fmt.Fprintf(w, "Result 4 = Name:%s, ID:%d, Height:%d\n", randomPokemon.Name, randomPokemon.ID, randomPokemon.Height)
		// }

		fmt.Fprintf(w, "####################\n")

		tenRandomPokemons, err := getTenRandomPokemon()
		if err != nil {
			log.Fatalf("Error getting random pokemon: %v", err)
		}
		for _, pokemon := range tenRandomPokemons {
			fmt.Fprintf(w, "Result 5 = Name: %s, ID: %d, Height: %d\n", pokemon.Name, pokemon.ID, pokemon.Height)
		}

		fmt.Fprintf(w, "####################\n")

		pokemonResponse, err := getPokemonByType("fire")
		if err != nil {
			log.Fatalf("Error getting pokemon by type: %v", err)
		}
		for _, pokemon := range pokemonResponse {
			fmt.Fprintf(w, "Result 6 = Name: %s, ID: %d\n", pokemon.Name, pokemon.ID)
		}

		// fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
