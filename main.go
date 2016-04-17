package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	home := os.Getenv("HOME")
	var maybetoken string

	f, err := os.Open(home + "/.guru/relogin_token")
	defer f.Close()
	if err != nil {
		fmt.Println("relogin token not found at " + home + "./guru/relogin_token")
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(f)
	reloginToken := strings.TrimSpace(string(b))
	fmt.Println(reloginToken)

	_, err = ioutil.ReadFile(home + "/.guru/token")
	if err == nil {
		//TODO: right now i need to force auth for the team id
		//maybetoken = strings.TrimSpace(string(t))
	}

	client := guru.NewClient(&guru.Config{ReloginToken: reloginToken, Token: maybetoken})
	results := client.GetFacts("mesos", "docker")

	for _, card := range results {
		fmt.Println(card.Title, card.Content)
	}

	for _, board := range client.GetBoards() {
		fmt.Println(board.Title, board.Description)
	}

	//card := client.CreateCard(guru.NewCard("test", "testerino"))
	//fmt.Println(card.Id)

	tagCategories := client.GetTagCategories()
	defaultCategory := tagCategories[0]

	var lastTag string

	for _, tag := range defaultCategory.Tags {
		lastTag = tag.Id
		fmt.Println(tag.Id + " " + tag.Value)
	}

	cards := client.CardByTags(lastTag)

	fmt.Println("cards matching")
	for _, x := range cards {
		fmt.Println(x.Id + " " + x.Title)
	}

}
