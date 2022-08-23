## Elementary http client

Simple wrapper aroung <a href="https://chandan-02.github.io/anime-facts-rest-api/">anime-facts-rest-api</a>

## Example

```Go
package main

import (
	"log"
	"math/rand"
	"time"

	client "github.com/energywork/simple-http-client/anime_facts"
)

const (
	timeout = 30 * time.Second
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	client, err := client.NewClient(timeout)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ANIME LIST:")

	anime, err := client.GetAnimeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range anime {
		log.Println(a.String())
	}

	selectedAnime := anime[rand.Intn(len(anime))]

	log.Printf("FACTS(%s):\n", selectedAnime.Name)

	facts, err := client.GetAnimeFacts(selectedAnime.Name)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range facts {
		log.Println(f.String())
	}

	selectedFact := facts[rand.Intn(len(facts))]

	log.Printf("SPECIFIC FACT(%s:%d):", selectedAnime.Name, selectedFact.ID)

	fact, err := client.GetSpecificFact(selectedAnime.Name, selectedFact.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fact.String())
}

```

