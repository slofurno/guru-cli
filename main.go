package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"io/ioutil"
	"os"
	"strings"
)

func joinTags(card *guru.Card) string {
	var tagNames []string
	for _, tag := range card.Tags {
		tagNames = append(tagNames, tag.Value)
	}
	return strings.Join(tagNames, ", ")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("not enough args")
		os.Exit(1)
	}

	client := initClient()

	switch args[0] {
	case "create-card":
		title := args[1]
		content := strings.Join(args[2:], " ")
		card := client.CreateCard(guru.NewCard(title, content))
		fmt.Printf("%s %s\n", card.Id, card.Title)
	case "get-card":
		id := args[1]
		card := client.GetCard(id)
		tags := joinTags(card)
		fmt.Printf("%-40s  %s \n\n%s", card.Title, tags, card.Content)
	case "add-tags":
		id := args[1]
		if len(args) < 2 {
			fmt.Println("missing tags")
			os.Exit(1)
		}
		tags := args[2:]
		client.AddTags(id, tags)
	default:
		cards := client.QueryCards(args...)
		for _, card := range cards {
			fmt.Printf("%s \n%s\n\n", card.Title, card.Content)
		}
	}
}

func initClient() *guru.Client {
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

	t, err := ioutil.ReadFile(home + "/.guru/token")
	if err == nil {
		maybetoken = strings.TrimSpace(string(t))
	}

	config := &guru.Config{ReloginToken: reloginToken, Token: maybetoken}
	client := guru.NewClient(config)
	//TODO: something
	team := client.GetTeam()
	client.Config.Team = team.Id

	return client
}
