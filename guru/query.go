package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *Client) CardByTags(tagIds ...string) []*Card {
	qr := &QueryRequest{Query: DefaultQuery()}
	expression := DefaultExpression()
	expression.Ids = tagIds
	qr.Query.NestedExpressions = []*Expression{expression}

	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)

	err := encoder.Encode(qr)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(buffer.Bytes()))

	//TODO: check for 204: no content (if no matches)
	res, err := s.makeRequest("POST", "https://api.getguru.com/api/v1/search/query", buffer)
	if err != nil {
		fmt.Println(err)
	}

	cards := []*Card{}
	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&cards)
	if err != nil {
		fmt.Println(err)
	}
	return cards
}

type QueryRequest struct {
	Query *Query `json:"query"`
	//ShowArchived nil `json:"showArchived"`
	//Sorts        nil `json:"sorts"`
}

type Expression struct {
	Type string   `json:"type"`
	Ids  []string `json:"ids"`
	Op   string   `json:"op"`
}

type Query struct {
	Op                string        `json:"op"`
	Type              string        `json:"type"`
	NestedExpressions []*Expression `json:"nestedExpressions"`
}

func DefaultExpression() *Expression {
	return &Expression{Op: "EXISTS", Type: "tag"}
}

func DefaultQuery() *Query {
	return &Query{Op: "AND", Type: "grouping"}
}
