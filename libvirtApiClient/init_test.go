package libvirtApiClient

import (
	"net/http"
)

type MockDoRequester struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockDoRequester) Do(req *http.Request) (*http.Response, error) {
	return m.MockResponse, m.MockError
}
