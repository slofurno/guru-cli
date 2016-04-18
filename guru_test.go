package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"testing"
)

func TestEverything(t *testing.T) {
	client := initClient()

	results := client.SearchCards("rick", "mesos")
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

	cards := client.QueryCards(lastTagId)

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
