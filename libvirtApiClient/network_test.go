package libvirtApiClient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func Test_Network_CreateNetwork(t *testing.T) {

	mockResponse := []byte(`{"name":"TestNode","id":1,"status":0}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	network, err := client.CreateNetwork("TestNode")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if network.Name != "TestNode" {
		t.Errorf("Expected network name to be 'TestNode', got '%v'", network.Name)
	}

}

func Test_Network_GetNetwork(t *testing.T) {

	mockResponse := []byte(`{"name":"TestNode","id":1,"status":0}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	network, err := client.GetNetwork(1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if network.Name != "TestNode" {
		t.Errorf("Expected network name to be 'TestNode', got '%v'", network.Name)
	}

}

func Test_Network_DeleteNetwork(t *testing.T) {

	mockResponse := []byte(`{"name":"TestNode","id":1,"status":0}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	err := client.DeleteNetwork(1)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

}

func Test_Network_GetNetworkByName(t *testing.T) {

	mockResponse := []byte(`{"name":"TestNode","id":1,"status":0}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	network, err := client.GetNetworkByName("TestNode")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if network.Name != "TestNode" {
		t.Errorf("Expected network name to be 'TestNode', got '%v'", network.Name)
	}
}

func Test_Network_UpdateNetwork(t *testing.T) {

	update_network := NetworkR{Name: "TestNode", ID: 1, Status: 0}
	mockResponse, _ := json.Marshal(update_network)
	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	network, err := client.UpdateNetwork(update_network)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if network.Name != "TestNode" {
		t.Errorf("Expected network name to be 'TestNode', got '%v'", network.Name)
	}
}
