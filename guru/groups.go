package guru

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Id   string `json:"id"` //same as user email
	Type string `json:"type"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *Client) GetTeamMembers() []*Member {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/teams/%s/members", s.Config.Team)
	res, _ := s.makeRequest("GET", uri, nil)
	members := []*Member{}
	err := json.NewDecoder(res.Body).Decode(&members)
	if err != nil {
		fmt.Println(err.Error())
	}

	return members
}

func (s *Client) GetGroups() []*Group {
	uri := fmt.Sprintf("https://api.getguru.com/api/v1/teams/%s/groups", s.Config.Team)
	res, _ := s.makeRequest("GET", uri, nil)
	groups := []*Group{}
	_ = json.NewDecoder(res.Body).Decode(&groups)
	return groups
}
