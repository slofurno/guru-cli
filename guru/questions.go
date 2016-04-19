package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *Client) GetQuestions() []*Card {
	uri := "https://api.getguru.com/api/v1/tasks"
	res, _ := s.makeRequest("GET", uri, nil)
	tasks := []*Card{}
	err := json.NewDecoder(res.Body).Decode(&tasks)
	if err != nil {
		fmt.Println(err.Error())
	}

	return tasks
}

type User struct {
	Email string `json:"email"`
}

type Question struct {
	Question  string    `json:"preferredPhrase"`
	Verifiers []*Expert `json:"verifiers"`
}

type Expert struct {
	Type string `json:"type"` //user or user-group
	//only one of these
	UserGroup *Group `json:"userGroup, omitempty"`
	User      *User  `json:"user, omitempty"`
}

//only need card title, content, + id
func (s *Client) AnswerQuestion(cardId, answer string) {
	card := s.GetCard(cardId)
	card.Content = answer
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/questions/%s/answer", card.Id)
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(card)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = s.makeRequest("PUT", uri, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *Client) AskQuestion(tr *Question) *Card {
	uri := "https://api.getguru.com/api/v1/questions"
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := s.makeRequest("POST", uri, buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	card := &Card{}
	err = json.NewDecoder(res.Body).Decode(card)
	if err != nil {
		fmt.Println(err.Error())
	}

	return card
}
