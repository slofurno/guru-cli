package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Card struct {
	Id                   string  `json:"id, omitempty"`
	Title                string  `json:"preferredPhrase"`
	Content              string  `json:"content"`
	VerificationInterval int     `json:"verificationInterval"`
	ShareStatus          string  `json:"shareStatus"`
	CardType             string  `json:"cardType"`
	Tags                 []*Tag  `json:"tags, omitempty"`
	FileProvider         *string `json:"fileProvider, omitempty"`
	FileLink             *string `json:"fileLink, omitempty"`
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

//https://api.getguru.com/api/v1/cards/
func (s *Client) CreateCard(card *Card) *Card {
	uri := "https://api.getguru.com/api/v1/cards/"
	body := bytes.NewBuffer(nil)
	err := json.NewEncoder(body).Encode(card)

	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := s.makeRequest("POST", uri, body)

	if err != nil {
		fmt.Println(err.Error())
	}

	ret := &Card{}
	err = json.NewDecoder(res.Body).Decode(ret)

	if err != nil {
		fmt.Println(err.Error())
	}

	return ret
}

func (s *Client) ArchiveCard(card *Card) {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/cards/%v", card.Id)
	res, _ := s.makeRequest("DELETE", uri, nil)

	fmt.Println(res.Status)
}

func (s *Client) SearchCards(query ...string) []*Card {
	qs := strings.Join(query, ",")

	uri := fmt.Sprintf("https://api.getguru.com/api/v1/search?terms=%v", qs)
	res, _ := s.makeRequest("GET", uri, nil)

	decoder := json.NewDecoder(res.Body)
	results := []*Card{}
	_ = decoder.Decode(&results)
	return results
}

func (s *Client) GetCard(id string) *Card {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/cards/%s", id)
	res, _ := s.makeRequest("GET", uri, nil)

	decoder := json.NewDecoder(res.Body)
	card := &Card{}
	_ = decoder.Decode(card)
	return card
}

func (s *Client) AddTags(cardId string, tags []string) {
	defaultCat := s.GetTagCategories()[0]
	tagMap := map[string]string{}
	for _, tag := range defaultCat.Tags {
		tagMap[tag.Value] = tag.Id
	}

	for _, tag := range tags {
		if tagId, ok := tagMap[tag]; ok {
			s.AddTag(cardId, tagId)
		}
	}
}

func (s *Client) AddTag(cardId, tagId string) {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/cards/%s/tags/%s", cardId, tagId)
	_, _ = s.makeRequest("PUT", uri, nil)
}
