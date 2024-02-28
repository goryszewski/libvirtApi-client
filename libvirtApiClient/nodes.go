package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetIPByNodeName(nodename string) (*Worker, error) {

	payload, err := json.Marshal(map[string]string{"function": "GetNodeByHostname", "hostname": nodename})
	if err != nil {
		return &Worker{}, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/node", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return &Worker{}, err
	}

	body, err := c.doRequest(request, &c.Token)
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
