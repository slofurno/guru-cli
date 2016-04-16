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
	config *Config
}

type Config struct {
	Token  string
	Cookie string
}

type Auth struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

func NewClient(config *Config) *Client {

	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.getguru.com/user/auth", nil)
	req.Header.Set("Cookie", config.Cookie)
	res, err := client.Do(req)

	if err != nil {
		os.Exit(1)
	}

	decoder := json.NewDecoder(res.Body)
	auth := &Auth{}
	err = decoder.Decode(auth)

	if err != nil {
		os.Exit(1)
	}

	token := fmt.Sprintf("%v:%v", auth.Email, auth.Token)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	authtoken := fmt.Sprintf("Basic %v", encoded)

	return &Client{
		Client: client,
		token:  authtoken,
		config: config,
	}
}

func (s *Client) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", s.token)
	return s.Do(req)
}

func reduce(acc string, xs []string, fn func(string, string) string) string {
	if len(xs) == 0 {
		return acc
	}

	i := 0
	if acc == "" {
		i = 1
		acc = xs[0]
	}

	for ; i < len(xs); i++ {
		acc = fn(acc, xs[i])
	}

	return acc
}
