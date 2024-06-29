package libvirtApiClient

import (
	"net/http"
)

const HostURL string = "http://127.0.0.1:8050"

type DoRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HostURL    string
	HTTPClient DoRequester
	Token      string
	Auth       AuthStruct
}

type Config struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Url      *string `json:"url"`
}

func NewClient(conf Config, requester DoRequester) (*Client, error) {
	c := Client{
		HTTPClient: requester,
		HostURL:    HostURL,
	}

	if conf.Url != nil {
		c.HostURL = *conf.Url
	}

	if conf.Username == nil || conf.Password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *conf.Username,
		Password: *conf.Password,
	}

	authRequest, err := c.SignIn()

	if err != nil {
		return nil, err
	}
	c.Token = authRequest.Token

	return &c, nil
}
