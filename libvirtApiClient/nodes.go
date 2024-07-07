package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetNodeByMetadata() (*NodeV2, error) {
	var url string = fmt.Sprintf("%s/api/v2/metadata", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("problem wih NewRequest: %v", err.Error())
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("problem with doRequest: %v", err.Error())
	}

	var node NodeV2

	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, fmt.Errorf("problem with Unmarshal: %v", err.Error())
	}

	return &node, nil

}

func (c *Client) GetNodeByName(name string) (*NodeV2, error) {

	var url string = fmt.Sprintf("%s/api/v2/node/%s", c.HostURL, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("problem wih NewRequest: %v", err.Error())
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("problem with doRequest: %v", err.Error())
	}

	var node NodeV2

	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, fmt.Errorf("problem with Unmarshal: %v", err.Error())
	}

	return &node, nil
}

func (c *Client) GetNodes() (*[]NodeV2, error) {
	var url string = fmt.Sprintf("%v/api/v2/node", c.HostURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("problem wih NewRequest: %v", err.Error())
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("problem with doRequest: %v", err.Error())
	}

	var nodes []NodeV2

	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, fmt.Errorf("problem with Unmarshal: %v", err.Error())
	}

	return &nodes, nil
}
