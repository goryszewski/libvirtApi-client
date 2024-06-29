package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

// GetIPByNodeName [deprecated]
func (c *Client) GetIPByNodeName(nodename string) (*Worker, error) { // DEP

	payload, err := json.Marshal(map[string]string{"function": "GetNodeByHostname", "hostname": nodename})
	if err != nil {
		return &Worker{}, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/node", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return &Worker{}, err
	}

	body, _, err := c.doRequest(request)
	if err != nil {
		return &Worker{}, err
	}

	var node Worker

	err = json.Unmarshal(body, &node)

	if err != nil {
		return &Worker{}, err
	}

	return &node, nil

}
