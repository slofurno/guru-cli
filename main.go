package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"os"
)

func main() {

	cookie := os.Getenv("GURU_COOKIE")
	client := guru.NewClient(&guru.Config{Cookie: cookie})

	results := client.GetFacts("mesos", "docker")

	for _, card := range results {
		fmt.Println(card.Title, card.Type)
	}

	for _, board := range client.GetBoards() {
		fmt.Println(board.Title, board.Description)
	}

	/*
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://api.getguru.com/api/v1/search", nil)
		req.Header.Set("Authorization", TOKEN)
		res, _ := client.Do(req)

		decoder := json.NewDecoder(res.Body)
		results := []*Fact{}

		err := decoder.Decode(&results)

		if err != nil {
			fmt.Println(err)
		}


		fmt.Println(results)
	*/

}
