package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestEverything(t *testing.T) {

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

	storedtoken, err := ioutil.ReadFile(home + "/.guru/token")
	if err == nil {
		//TODO: right now i need to force auth for the team id
		maybetoken = strings.TrimSpace(string(storedtoken))
	}

	client := guru.NewClient(&guru.Config{ReloginToken: reloginToken, Token: maybetoken})
	team := client.GetTeam()
	client.Config.Team = team.Id

	results := client.SearchCards("asdasd", "sadasdas")
	cardIds := []string{}

	fmt.Println("search results:")
	for _, card := range results {
		cardIds = append(cardIds, card.Id)
		fmt.Println(card.Title, card.Content)
	}

	fmt.Println("board titles / descriptions:")
	for _, board := range client.GetBoards() {
		fmt.Println(board.Title, board.Description)
	}

	//card := client.CreateCard(guru.NewCard("test", "testerino"))
	//fmt.Println(card.Id)

	tagCategories := client.GetTagCategories()
	defaultCategory := tagCategories[0]

	var lastTagId string

	fmt.Println("all of our tag ids + values")
	for _, tag := range defaultCategory.Tags {
		lastTagId = tag.Id
		fmt.Println(tag.Id + " " + tag.Value)
	}

	cards := client.CardByTags(lastTagId)

	fmt.Printf("\ncards with the tag %s: \n", lastTagId)
	for _, x := range cards {
		fmt.Println(x.Id + " " + x.Title)
	}

	fmt.Println(fmt.Sprintf("adding tag to %v cards", len(cardIds)))

	_ = client.CreateTag(&guru.CreateTagRequest{
		CategoryId: defaultCategory.Id,
		Value:      "whatever2",
	})

	client.AddTagToCards(&guru.BulkRequest{
		&guru.BulkAction{
			Type:   "tag-card",
			TagIds: []string{lastTagId},
		},
		&guru.BulkItems{
			Type:    "id",
			CardIds: cardIds,
		},
	})
}