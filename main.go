package main

import (
	"fmt"
	"github.com/slofurno/guru-cli/guru"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("not enough args")
		os.Exit(1)
	}

	client := initClient()

	switch args[0] {
	case "find":
		defaultCat := client.GetTagCategories()[0]
		tagMap := map[string]string{}
		for _, tag := range defaultCat.Tags {
			tagMap[tag.Value] = tag.Id
		}

		tags := []string{}
		for _, tag := range strings.Split(args[1], ",") {
			tags = append(tags, tagMap[tag])
		}

		cards := client.CardByTags(tags...)
		for _, card := range cards {
			fmt.Printf("%-22s %s\n", card.Id, card.Title)
		}

	case "create-card":
		title := args[1]
		content := strings.Join(args[2:], " ")
		client.CreateCard(guru.NewCard(title, content))
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
