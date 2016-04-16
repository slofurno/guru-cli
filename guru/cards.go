package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Card struct {
	Id                   string `json:"id, omitempty"`
	Title                string `json:"preferredPhrase"`
	Content              string `json:"content"`
	VerificationInterval int    `json:"verificationInterval"`
	ShareStatus          string `json:"shareStatus"`
	CardType             string `json:"cardType"`
}

func NewCard(title, content string) *Card {
	return &Card{
		Title:                title,
		Content:              content,
		VerificationInterval: 30,
		ShareStatus:          "TEAM",
		CardType:             "CARD",
	}
}

func (s *Client) UpdateCard(card *Card) int {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/cards/%v", card.Id)
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.Encode(card)

	res, _ := s.makeRequest("PUT", uri, buffer)

	return res.StatusCode
}

func (s *Client) CreateCard(card *Card) *Card {
	uri := "https://api.getguru.com/api/v1/cards/"
	body := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(body)
	_ = encoder.Encode(card)

	res, err := s.makeRequest("POST", uri, body)

	if err != nil {
		fmt.Println(err.Error())
	}

	ret := &Card{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(ret)

	if err != nil {
		fmt.Println(err.Error())
	}

	return ret
}

func (s *Client) GetFacts(query ...string) []*Card {
	qs := reduce("", query, func(a string, c string) string {
		return a + "," + c
	})

	uri := fmt.Sprintf("https://api.getguru.com/api/v1/search?terms=%v", qs)
	res, _ := s.makeRequest("GET", uri, nil)

	decoder := json.NewDecoder(res.Body)
	results := []*Card{}
	_ = decoder.Decode(&results)
	return results
}
