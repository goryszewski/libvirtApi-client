package libvirtApiClient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func Test_Nodes_GetIPByNodeName(t *testing.T) {

	mockResponse := []byte(`{"Name":"TestNode","Internal":"192.168.1.1","External":"8.8.8.8","Type":"ExampleType"}`)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponse)),
		Header:     make(http.Header),
	}

	cf := Config{}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, _ := NewClient(cf, requester)

	worker, err := client.GetNodeByName("TestNode")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if worker.Name != "TestNode" {
		t.Errorf("Expected worker name to be 'TestNode', got '%v'", worker.Name)
	}

	// if worker.Interface[0].Address != "192.168.1.1" {
	// 	t.Errorf("Expected IP.Internal to be '192.168.1.1', got '%v'", worker.Interface[0].Address)
	// }

	// if worker.Interface[1].Address != "8.8.8.8" {
	// 	t.Errorf("Expected IP.External to be '8.8.8.8', got '%v'", worker.Interface[1].Address)
	// }

}
