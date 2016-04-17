package guru

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tags []*Tag `json:"tags"`
}

func (s *Client) GetTagCategories() []*Category {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/teams/%v/tagcategories/", s.config.Team)
	res, _ := s.makeRequest("GET", uri, nil)
	decoder := json.NewDecoder(res.Body)
	tagCategories := []*Category{}
	err := decoder.Decode(&tagCategories)

	if err != nil {
		fmt.Println(err.Error())
	}

	return tagCategories
}
