package libvirtApiClient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

var test_payload ServiceLoadBalancer = ServiceLoadBalancer{
	Name:      "nnn21",
	Namespace: "nnn",
	Ports: []Port_Service{
		Port_Service{
			Name:     "test",
			Protocol: "TCP",
			Port:     1,
			NodePort: 2,
		},
	},
	Nodes: []Node{
		Node{
			Name:       "test",
			Private_ip: "10.10.11.1",
			Public_ip:  "192.168.1.1",
		},
	},
}
var test_response ServiceLoadBalancerRespons = ServiceLoadBalancerRespons{
	ID:                  "1",
	Ip:                  "10.10.10.1",
	ServiceLoadBalancer: &test_payload,
}

func Test_k8s_GetLoadBalancer(t *testing.T) {

	mockResponse, _ := json.Marshal(test_response)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	ip, err := client.GetLoadBalancer(test_payload)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ip != test_response.Ip {
		t.Errorf("Expected lb ip to be '10.10.10.1', got '%v'", ip)
	}

}
