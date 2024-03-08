package libvirtApiClient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockDoRequester struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockDoRequester) Do(req *http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

func Test_Nodes_GetIPByNodeName(t *testing.T) {

	mockResponse := []byte(`{"Name":"TestNode","IP":{"Private":"192.168.1.1","Public":"8.8.8.8"},"Type":"ExampleType"}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	worker, err := client.GetIPByNodeName("TestNode")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if worker.Name != "TestNode" {
		t.Errorf("Expected worker name to be 'TestNode', got '%v'", worker.Name)
	}

	if worker.IP.Private != "192.168.1.1" {
		t.Errorf("Expected IP.Private to be '192.168.1.1', got '%v'", worker.IP.Private)
	}

	if worker.IP.Public != "8.8.8.8" {
		t.Errorf("Expected IP.Public to be '8.8.8.8', got '%v'", worker.IP.Public)
	}

}
