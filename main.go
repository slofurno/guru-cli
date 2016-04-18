package main

import (
	"encoding/json"
	"errors"
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

func getLogin() (*guru.Login, error) {
	maybeEmail := os.Getenv("GURU_EMAIL")
	maybePass := os.Getenv("GURU_PASS")
	home := os.Getenv("HOME")

	if maybeEmail != "" || maybePass != "" {
		return &guru.Login{Email: maybeEmail, Password: maybePass}, nil
	}

	f, err := os.Open(home + "/.guru/credentials")
	defer f.Close()
	if err != nil {
		return nil, errors.New("credentials not set in env or $HOME/.guru/credentials")
	}

	b, _ := ioutil.ReadAll(f)
	login := &guru.Login{}
	err = json.Unmarshal(b, login)

	if err != nil {
		return nil, errors.New("credentials file invalid format")
	}

	return login, nil
}

//TODO: return error and exit in one place
func initClient() *guru.Client {
	config := &guru.Config{}
	home := os.Getenv("HOME")
	var client *guru.Client

	t, err := ioutil.ReadFile(home + "/.guru/token")
	if err == nil {
		config.Token = strings.TrimSpace(string(t))
	}

	if reloginToken, err := getReloginToken(); err == nil {
		config.ReloginToken = reloginToken
		client = guru.NewClient(config)
	} else if login, err := getLogin(); err == nil {
		client = guru.NewClient(config)
		if err = client.Login(login); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		f, _ := os.Create(home + "/.guru/relogin_token")
		defer f.Close()
		f.WriteString(fmt.Sprintf("%s\n", config.ReloginToken))

	} else {
		fmt.Println("neither login credentials nor relogin token found")
		os.Exit(1)
	}

	//TODO: something
	team := client.GetTeam()
	client.Config.Team = team.Id

	return client
}

func getReloginToken() (string, error) {
	home := os.Getenv("HOME")
	f, err := os.Open(home + "/.guru/relogin_token")
	defer f.Close()
	if err != nil {
		return "", errors.New("relogin token not found at " + home + "./guru/relogin_token")
	}

	b, _ := ioutil.ReadAll(f)
	return strings.TrimSpace(string(b)), nil

}
