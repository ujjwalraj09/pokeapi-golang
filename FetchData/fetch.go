package api

import (
	"log"
	"net/http"
)

func Url(pokemonName string) string {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	return url
}

func FetchUrl(url string) http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	return *resp
}
