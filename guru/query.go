package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *Client) QueryCards(tags ...string) []*Card {
	defaultCat := s.GetTagCategories()[0]
	tagMap := map[string]string{}
	for _, tag := range defaultCat.Tags {
		tagMap[tag.Value] = tag.Id
	}

	expressions := []*Expression{{CardType: "QUESTION", Op: "NE", Type: "card-type"}}
	for _, tag := range tags {
		if tagId, ok := tagMap[tag]; ok {
			expressions = append(expressions,
				&Expression{Type: "tag", Op: "EXISTS", Ids: []string{tagId}})
		}
	}

	query := &QueryRequest{
		Query: &Query{
			NestedExpressions: expressions,
			Op:                "AND",
			Type:              "grouping",
		},
		SearchTerms: "",
	}

	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(query)

	if err != nil {
		fmt.Println(err.Error())
	}
	//TODO: check for 204: no content (if no matches)
	//FIXME: this won't populate tags :(
	res, err := s.makeRequest("POST", "https://api.getguru.com/api/v1/search/query", buffer)
	if err != nil {
		fmt.Println(err)
	}

	cards := []*Card{}
	err = json.NewDecoder(res.Body).Decode(&cards)

	if err != nil {
		fmt.Println(err)
	}
	return cards
}

type QueryRequest struct {
	Query       *Query `json:"query"`
	SearchTerms string `json:"searchTerms"`
	//ShowArchived nil `json:"showArchived"`
	//Sorts        nil `json:"sorts"`
}

type Expression struct {
	Type     string   `json:"type"`
	Ids      []string `json:"ids, omitempty"`
	Op       string   `json:"op"`
	CardType string   `json:"cardType, omitempty"`
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
