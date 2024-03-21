package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetLoadBalancer(bind_payload ServiceLoadBalancer) (string, error) {

	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/lb/%v/%v", c.HostURL, bind_payload.Namespace, bind_payload.Name), nil)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer NewRequest (%v)", err)
		return "", err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer doRequest (%v)", err)
		return "", err
	}

	var lb ServiceLoadBalancerRespons

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer Unmarshal (%v)", err)
		return "", err
	}

	return lb.Ip, nil
}

func (c *Client) CreateLoadBalancer(bind_payload ServiceLoadBalancer) (string, error) {

	payload, err := json.Marshal(bind_payload)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer Marshal (%v)", err)
		return "", err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer NewRequest (%v)", err)
		return "", err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer doRequest (%v)", err)
		return "", err
	}

	var lb ServiceLoadBalancerRespons

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error CreateLoadBalancer Unmarshal (%v)", err)
		return "", err
	}

	return lb.Ip, nil
}

func (c *Client) DeleteLoadBalancer(bind_payload ServiceLoadBalancer) error {
	payload, err := json.Marshal(bind_payload)
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer Marshal (%v)", err)
		return err
	}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer NewRequest (%v)", err)
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer doRequest (%v)", err)
		return err
	}

	var lb ServiceLoadBalancerRespons

	err = json.Unmarshal(body, &lb)

	if err != nil {
		log.Fatalf("Error DeleteLoadBalancer (%v)", err)
		return err
	}
	return nil

}

func (c *Client) UpdateLoadBalancer(bind_payload ServiceLoadBalancer) error {

	payload, err := json.Marshal(bind_payload)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer Marshal (%v)", err)
		return err
	}

	request, err := http.NewRequest("PUT", fmt.Sprintf("%v/api/lb", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer NewRequest (%v)", err)
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer doRequest (%v)", err)
		return err
	}

	var lb ServiceLoadBalancerRespons

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer Unmarshal (%v)", err)
		return err
	}

	return nil

}
