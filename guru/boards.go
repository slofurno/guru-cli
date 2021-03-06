package guru

import (
	"encoding/json"
	"fmt"
)

type Board struct {
	Id          string  `json:"id, omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Items       []*Card `json:"items"`
}

func (s *Client) GetBoards() []*Board {
	url := "https://api.getguru.com/api/v1/boards"
	res, _ := s.makeRequest("GET", url, nil)

	boards := []*Board{}
	err := json.NewDecoder(res.Body).Decode(&boards)

	if err != nil {
		fmt.Println(err.Error())
	}

	return boards
}
