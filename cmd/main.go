package main

import (
	"fmt"
	"os"
	api "pokeapi/FetchData"
	jsonparse "pokeapi/JsonParse"
	saveimage "pokeapi/Saveimage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <Pokemon Name>")
		return
	}
	pokemonName := os.Args[1]
	resp := api.FetchUrl(api.Url(pokemonName))
	pokemon := jsonparse.ParsePokemonJson(resp)

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("National Number: %d\n", pokemon.Id)

	fmt.Print("Type:")
	for _, typevar := range pokemon.Types {
		fmt.Printf("%s ", typevar.Type.Name)
	}
	fmt.Println()

	fmt.Print("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("%s ", ability.Ability.Name)
	}
	fmt.Println()

	fmt.Printf("Height: %0.1f\n", float32(pokemon.Height)/10)
	fmt.Printf("Weight: %0.1f\n", float32(pokemon.Weight)/10)
	fmt.Println()
	fmt.Println("Base Stats: ")
	sum := 0
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s: %d \n", stat.Stat.Name, stat.BaseStat)
		sum += stat.BaseStat

	}
	fmt.Println("Total: ", sum)

	imageUrl := pokemon.Sprites.ImageUrl
	saveimage.SaveImageFromURL(imageUrl, "sprite.png")

}
