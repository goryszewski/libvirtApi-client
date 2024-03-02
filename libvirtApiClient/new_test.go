package libvirtApiClient

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	cf := Config{}
	requester := &http.Client{Timeout: 10 * time.Second}
	client, err := NewClient(cf, requester)
	if client == nil && err != nil {
		t.Fatalf(`%v = %v`, client, err)
	}
}

func TestNewClient_URL(t *testing.T) {
	cf := Config{}
	requester := &http.Client{Timeout: 10 * time.Second}
	client, err := NewClient(cf, requester)
	if client.HostURL != "http://127.0.0.1:8050" {
		t.Fatalf(`%v = %v`, client.HostURL, err)
	}
}
