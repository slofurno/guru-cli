package guru

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	*http.Client
	token  string
	Config *Config
}

type Config struct {
	Token        string
	ReloginToken string
	Team         string
}

type Auth struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Team  *Team  `json:"team"`
}

type Team struct {
	Id string `json:"id"`
}

func NewClient(config *Config) *Client {
	return &Client{
		Client: &http.Client{},
		token:  config.Token,
		Config: config,
	}
}

func (s *Client) auth() {
	req, _ := http.NewRequest("POST", "https://api.getguru.com/user/auth", nil)
	req.Header.Set("Cookie", s.Config.ReloginToken)
	res, err := s.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	decoder := json.NewDecoder(res.Body)
	auth := &Auth{}
	err = decoder.Decode(auth)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	token := fmt.Sprintf("%v:%v", auth.Email, auth.Token)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	authtoken := fmt.Sprintf("Basic %v", encoded)

	fmt.Println("authed and set token:", authtoken)
	home := os.Getenv("HOME")

	f, err := os.Create(home + "/.guru/token")
	defer f.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	f.WriteString(fmt.Sprintf("%v\n", authtoken))
	s.Config.Team = auth.Team.Id
	s.token = authtoken
}

func (s *Client) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	for {
		req, _ := http.NewRequest(method, url, body)
		req.Header.Set("Authorization", s.token)

		if method == "POST" || method == "PUT" {
			req.Header.Set("Content-Type", "application/json")
		}

		res, err := s.Do(req)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if res.StatusCode == http.StatusUnauthorized {
			fmt.Println("getting new token")
			s.auth()
		} else {
			return res, err
		}
	}
}
