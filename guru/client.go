package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	*http.Client
	token string
}

type Config struct {
	Token string
}

func NewClient(config *Config) *Client {
	return &Client{
		Client: &http.Client{},
		token:  config.Token,
	}
}

func reduce(acc string, xs []string, fn func(string, string) string) string {
	if len(xs) == 0 {
		return acc
	}

	i := 0
	if acc == "" {
		i = 1
		acc = xs[0]
	}

	for ; i < len(xs); i++ {
		acc = fn(acc, xs[i])
	}

	return acc
}

func (s *Client) UpdateCard(card *Card) int {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/cards/%v", card.Id)
	req, _ := http.NewRequest("PUT", uri, nil)
	req.Header.Set("Authorization", s.token)
	res, _ := s.Do(req)

	return res.StatusCode
}

func (s *Client) CreateCard(card *Card) *Card {
	uri := "https://api.getguru.com/api/v1/cards/"
	body := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(body)
	_ = encoder.Encode(card)
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Authorization", s.token)

	res, err := s.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	ret := &Card{}
	decoder := json.NewDecoder(res.Body)
	_ = decoder.Decode(ret)

	return ret
}

func (s *Client) GetFacts(query ...string) []*Fact {
	qs := reduce("", query, func(a string, c string) string {
		return a + "," + c
	})

	uri := fmt.Sprintf("https://api.getguru.com/api/v1/search?terms=%v", qs)
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("Authorization", s.token)
	res, _ := s.Do(req)

	decoder := json.NewDecoder(res.Body)
	results := []*Fact{}
	_ = decoder.Decode(&results)
	return results
}
