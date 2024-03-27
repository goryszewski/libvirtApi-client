package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetLoadBalancer(bind_payload ServiceLoadBalancer) (*ServiceLoadBalancerResponse, bool, error) {

	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/lb/%v/%v", c.HostURL, bind_payload.Namespace, bind_payload.Name), nil)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer NewRequest (%v)", err)
		return nil, false, err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer doRequest (%v)", err)
		return nil, false, err
	}

	var lb ServiceLoadBalancerResponse

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error GetLoadBalancer Unmarshal (%v)", err)
		return nil, false, err
	}
	if lb.Ip == "" { // DOTO change to http code 404
		return nil, false, nil
	}

	return &lb, true, nil
}

func (c *Client) GetAllLoadBalancers() ([]ServiceLoadBalancerResponse, error) {

	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/lb", c.HostURL), nil)
	if err != nil {
		log.Fatalf("Error GetAllLoadBalancers NewRequest (%v)", err)
		return nil, err
	}

	body, err := c.doRequest(request)
	if err != nil {
		log.Fatalf("Error GetAllLoadBalancers doRequest (%v)", err)
		return nil, err
	}

	var lbs []ServiceLoadBalancerResponse

	err = json.Unmarshal(body, &lbs)
	if err != nil {
		log.Fatalf("Error GetAllLoadBalancers Unmarshal (%v)", err)
		return nil, err
	}

	return lbs, nil
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

	var lb ServiceLoadBalancerResponse

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

	var lb ServiceLoadBalancerResponse

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

	var lb ServiceLoadBalancerResponse

	err = json.Unmarshal(body, &lb)
	if err != nil {
		log.Fatalf("Error UpdateLoadBalancer Unmarshal (%v)", err)
		return err
	}

	return nil

}
