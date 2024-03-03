package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type NetworkR struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (n *Client) GetNetwork(id types.Int64) (*NetworkR, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/network/%v", *n.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	body, err := n.doRequest(request)
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

func (n *Client) DeleteNetwork(id types.Int64) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/network/%v", *n.HostURL, id), nil)
	if err != nil {
		return err
	}
	body, err := n.doRequest(request)
	fmt.Print(body)
	if err != nil {
		return err
	}
	return nil
}

func (n *Client) CreateNetwork(name string) (*NetworkR, error) {
	payload, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/network", *n.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := n.doRequest(request)
	if err != nil {
		return nil, err
	}
	var net NetworkR

	err = json.Unmarshal(body, &net)

	if err != nil {
		log.Fatal("Error CreateNetwork", err)
	}
	return &net, nil
}