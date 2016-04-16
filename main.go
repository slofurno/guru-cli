package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
)

const TOKEN = "Basic c2xvZnVybm9AZ21haWwuY29tOjBkZjQzN2QwLWY1NjgtNGViOS04MTMyLTk4MjVkZjdkMmJhOA=="

func main() {

	client := guru.NewClient(&guru.Config{Token: TOKEN})

	results := client.GetFacts("mesos", "docker")

	for _, card := range results {
		fmt.Println(card.Title, card.Type)
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
