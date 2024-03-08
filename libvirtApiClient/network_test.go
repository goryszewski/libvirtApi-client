package libvirtApiClient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockDoRequester_network struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockDoRequester_network) Do(req *http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

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
