package guru

import (
	"encoding/json"
	"fmt"
)

func (s *Client) GetTeam() *Team {
	uri := "https://api.getguru.com/api/v1/teams"
	res, _ := s.makeRequest("GET", uri, nil)
	teams := []*Team{}
	err := json.NewDecoder(res.Body).Decode(&teams)

	if err != nil {
		fmt.Println(err.Error())
	}
	//TODO: lets assume one team for now
	return teams[0]
}
