package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type RedisEntry struct {
	Key   string
	Value Stat
}

type Stat struct {
	HP      int // base hp
	Attack  int // base attack
	Defense int // base defense
	Speed   int // base speed
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Context for the operations
	ctx := context.Background()

	// Get all keys
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}

	var redisData []RedisEntry
	for _, key := range keys {
		value, err := client.Get(ctx, key).Result()
		if err != nil {
			fmt.Printf("Error getting value for key %s: %s\n", key, err)
		} else {
			var stat Stat
			err := json.Unmarshal([]byte(value), &stat)
			if err != nil {
				fmt.Printf("Error unmarshalling JSON for key %s: %s\n", key, err)
			} else {
				redisData = append(redisData, RedisEntry{Key: key, Value: stat})
			}
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Use redisData here
		tmpl, err := template.ParseFiles("./index.html")
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, redisData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)

}
