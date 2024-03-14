package libvirtApiClient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func Test_Auth_GetToken(t *testing.T) {
	mockresponse := []byte(`{"user_id":1,"username":"2","token":"3"}`)
	Username := "test"
	Password := "test"
	URL := "https://128.0.0.1:8050"

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockresponse)),
		Header:     make(http.Header),
	}

	cf := Config{&Username, &Password, &URL}
	requester := &MockDoRequester{MockResponse: mockHttpResponse, MockError: nil}
	client, err := NewClient(cf, requester)
	if client == nil {
		t.Errorf("Expected not nil '%v'", err)
	}
	token := client.GetToken()
	if token != "3" {
		t.Errorf("Expected token 3 not'%v'", token)
	}
}
