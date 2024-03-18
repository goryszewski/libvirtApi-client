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
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), nil)
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
func (c *Client) CreateLoadBalancer(bind_payload bindServiceLB) (string, error) {

	payload, err := json.Marshal(bind_payload)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer Marshal (%v)", err)
		return "", err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer NewRequest (%v)", err)
		return "", err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer doRequest (%v)", err)
		return "", err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer Unmarshal (%v)", err)
		return "", err
	}

	return lb.Ip, nil
}

func (c *Client) DeleteLoadBalancer(service string) error {
	payload, err := json.Marshal(map[string]string{"service": service})
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer Marshal (%v)", err)
		return err
	}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer NewRequest (%v)", err)
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer doRequest (%v)", err)
		return err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)

	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer (%v)", err)
		return err
	}
	return nil

}

func (c *Client) UpdateLoadBalancer(bind_payload bindServiceLB) error {

	payload, err := json.Marshal(bind_payload)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer Marshal (%v)", err)
		return err
	}

	request, err := http.NewRequest("PUT", fmt.Sprintf("%v/api/v1/k8s/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer NewRequest (%v)", err)
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer doRequest (%v)", err)
		return err
	}

	var lb LB

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer Unmarshal (%v)", err)
		return err
	}

	return nil

}
