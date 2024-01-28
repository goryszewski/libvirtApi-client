package libvirtApiClient

import (
	"net/http"
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

func NewClient(hostname, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	return &c, nil
}
