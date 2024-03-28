package libvirtApiClient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

var data LoadBalancer = LoadBalancer{
	Name:      "nnn21",
	Namespace: "nnn",
	Ports: []Port_Service{
		Port_Service{
			Name:     "test",
			Protocol: "TCP",
			Port:     80,
			NodePort: 32011,
		},
	},
	Nodes: []Node{
		Node{
			Name:     "test",
			Internal: "10.10.11.1",
			External: "192.168.1.1",
		},
	},
	Ip: "10.10.10.1",
}

func Test_k8s_GetLoadBalancer(t *testing.T) {

	mockResponse, _ := json.Marshal(data)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	loadbalancer, is_exist, err := client.GetLoadBalancer(data)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !is_exist {
		t.Errorf("LB not exist")
	}

	if loadbalancer.Ip != data.Ip {
		t.Errorf("Expected lb ip to be '%v', got '%v'", data.Ip, loadbalancer.Ip)
	}

}
