package libvirtApiClient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockDoRequester_k8s struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockDoRequester_k8s) Do(req *http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

func Test_k8s_GetFreeLB(t *testing.T) {

	mockResponse := []byte(`{"id":1,"ip":"10.10.10.1","service":""}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	network, err := client.GetFreeLB()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if network.Ip != "10.10.10.1" {
		t.Errorf("Expected lb ip to be '10.10.10.1', got '%v'", network.Ip)
	}

}
