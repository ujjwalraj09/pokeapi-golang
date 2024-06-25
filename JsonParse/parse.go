package jsonparse

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	p "pokeapi/models"
)

func ParsePokemonJson(resp http.Response) p.Pokemon {

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	var pokemon p.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}
	return pokemon
}
