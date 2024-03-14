package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GET free ip /api/v1/k8s/lb

func (c *Client) GetFreeLB() (*LB, error) {
	payload, err := json.Marshal(map[string]string{"function": "GetFirstFreeLB"})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)

	if err != nil {
		log.Fatal("Error 003")
	}

	return &lb, nil
}
func (c *Client) BindLB(ip string, service string, workers string) error {

	payload, err := json.Marshal(map[string]string{"function": "BindLB", "ip": ip, "service": service, "workers": workers})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		return err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatal("Error 003")
	}

	return nil
}

func (c *Client) UnBindLB(service string) error {
	payload, err := json.Marshal(map[string]string{"function": "UnBindLB", "service": service})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		return err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)

	if err != nil {
		log.Fatal("Error 003")
	}
	return nil

}

func (c *Client) UpdateBind(service, nodes string) error {

	return nil

}
