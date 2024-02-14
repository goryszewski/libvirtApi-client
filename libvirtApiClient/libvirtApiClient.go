package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const HostURL string = "http://127.0.0.1:8050"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("Sprawdz login i haslo")
	}

	requestBody, err := json.Marshal(c.Auth)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%/api/v1/auth", c.HostURL), strings.NewReader(string(requestBody)))
	body, err := c.doRequest(request, nil)
	if err != nil {
		return nil, err
	}
	authResponse := AuthResponse{}

	err = json.Unmarshal(body, &authResponse)
	return &authResponse, nil

}

func NewClient(hostname, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if hostname != nil {
		c.HostURL = *hostname
	}

	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	authRequest, err := c.SignIn()

	if err != nil {
		return nil, err
	}
	c.Token = authRequest.Token

	return &c, nil
}

func (c *Client) doRequest(request *http.Request, authToken *string) ([]byte, error) {
	token := c.Token
	if authToken != nil {
		token = *authToken
	}

	request.Header.Set("Authorizaton", token)
	response, err := c.HTTPClient.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, Body: %s", response.StatusCode, body)
	}

	return body, nil
}
