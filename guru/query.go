package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func (s *Client) QueryCards(tags ...string) []*Card {
	searchTerms := strings.Join(tags, " ")
	query := &QueryRequest{
		Query: &Query{
			NestedExpressions: []*Expression{&Expression{
				CardType: "QUESTION",
				Op:       "NE",
				Type:     "card-type",
			}},
			Op:   "AND",
			Type: "grouping",
		},
		SearchTerms: searchTerms,
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
	CardType string   `json:"cardType"`
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
