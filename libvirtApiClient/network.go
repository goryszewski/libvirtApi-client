package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type NetworkR struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (c *Client) GetNetwork(id int) (*NetworkR, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/network/%v", c.HostURL, id), nil)

	if err != nil {
		return nil, err
	}
	body, _, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	var net NetworkR

	err = json.Unmarshal(body, &net)

	if err != nil {
		log.Fatal("Error GetNetwork", err)
	}
	return &net, nil
}

func (c *Client) GetNetworkByName(name string) (*NetworkR, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/network/name/%v", c.HostURL, name), nil)
	if err != nil {
		return nil, err
	}
	body, _, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	var net NetworkR

	err = json.Unmarshal(body, &net)

	if err != nil {
		log.Fatal("Error GetNetwork", err)
	}
	return &net, nil

}

func (c *Client) DeleteNetwork(id int) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/network/%v", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	_, _, err = c.doRequest(request)

	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateNetwork(network NetworkR) (*NetworkR, error) {
	payload, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PUT", fmt.Sprintf("%v/api/network", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	var net NetworkR

	err = json.Unmarshal(body, &net)

	if err != nil {
		log.Fatal("Error Unmarshal CreateNetwork: ", err)
	}
	return &net, nil
}

func (c *Client) CreateNetwork(name string) (*NetworkR, error) {
	payload, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/network", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	var net NetworkR

	err = json.Unmarshal(body, &net)

	if err != nil {
		log.Fatal("Error Unmarshal CreateNetwork: ", err)
	}
	return &net, nil
}
