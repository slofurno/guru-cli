package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type CreateTagRequest struct {
	CategoryId string `json:"categoryId"`
	Value      string `json:"value"`
}

func (s *Client) GetTagCategories() []*Category {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/teams/%v/tagcategories/", s.Config.Team)
	res, _ := s.makeRequest("GET", uri, nil)
	tagCategories := []*Category{}
	err := json.NewDecoder(res.Body).Decode(&tagCategories)

	if err != nil {
		fmt.Println(err.Error())
	}

	return tagCategories
}

func (s *Client) CreateTag(cr *CreateTagRequest) *Tag {
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(cr)

	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(string(buffer.Bytes()))

	uri := fmt.Sprintf("https://api.getguru.com/api/v1/teams/%v/tagcategories/tags", s.Config.Team)
	res, _ := s.makeRequest("POST", uri, buffer)

	//TODO: check for 400 status: tag already used
	tag := &Tag{}
	_ = json.NewDecoder(res.Body).Decode(tag)
	return tag
}

type BulkAction struct {
	Type   string   `json:"type"`
	TagIds []string `json:"tagIds"`
}

type BulkItems struct {
	Type    string   `json:"type"`
	CardIds []string `json:"cardIds"`
}

type BulkRequest struct {
	Action *BulkAction `json:"action"`
	Items  *BulkItems  `json:"items"`
}

func (s *Client) AddTagToCards(request *BulkRequest) {
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(request)

	if err != nil {
		fmt.Println(err.Error())
	}

	res, _ := s.makeRequest("POST", "https://api.getguru.com/api/v1/cards/bulkop", buffer)
	fmt.Println("bulkops status: " + res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

//DELETE https://api.getguru.com/api/v1/teams/f390146e-ebe5-42b3-b077-a632d5564789/tagcategories/tags/b8ef1e93-b4a5-4139-a1b1-af344b118fa7
