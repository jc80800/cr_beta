package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/redis/go-redis/v9"
)

const (
	MAX_POKEMON   = 2000
	HP_INDEX      = 0
	ATTACK_INDEX  = 1
	DEFENSE_INDEX = 2
	SPEED_INDEX   = 5
)

func log(name string, stat string) {
	fmt.Println(fmt.Sprintf("%s with %s has been inserted into db", name, stat))

}

func process(client *redis.Client, pokemon structs.Pokemon, ctx context.Context, pokemon_name string) {
	stats := pokemon.Stats

	hp := stats[HP_INDEX]
	attack := stats[ATTACK_INDEX]
	defense := stats[DEFENSE_INDEX]
	speed := stats[SPEED_INDEX]

	stat := Stat{HP: hp.BaseStat, Attack: attack.BaseStat, Defense: defense.BaseStat, Speed: speed.BaseStat}

	jsonData, err := json.Marshal(stat)

	if err != nil {
		panic(err)
	}

	err = client.Set(ctx, pokemon_name, string(jsonData), 0).Err()
	if err != nil {
		panic(err)
	}

	log(pokemon_name, string(jsonData))

}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Fetch the list of all Pokémon
	pokemons, err := pokeapi.Resource("pokemon", 0, MAX_POKEMON)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	// Iterate over each Pokémon and store its name and height in Redis
	for _, p := range pokemons.Results {
		pokemon, err := pokeapi.Pokemon(p.Name)
		if err != nil {
			fmt.Printf("Error fetching details for %s: %v\n", p.Name, err)
			continue
		}

		process(client, pokemon, ctx, p.Name)
	}
}
