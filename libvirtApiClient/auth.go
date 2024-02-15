package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("sprawdz login i haslo")
	}

	requestBody, err := json.Marshal(c.Auth)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/auth", c.HostURL), strings.NewReader(string(requestBody)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(request, nil)
	if err != nil {
		return nil, err
	}
	authResponse := AuthResponse{}

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}
	return &authResponse, nil

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
