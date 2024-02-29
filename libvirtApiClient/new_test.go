package libvirtApiClient

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	cf := Config{}
	client, err := NewClient(cf)
	if client == nil && err != nil {
		t.Fatalf(`%v = %v`, client, err)
	}
}

func TestNewClient_URL(t *testing.T) {
	cf := Config{}
	client, err := NewClient(cf)
	if client.HostURL != "http://127.0.0.1:8050" {
		t.Fatalf(`%v = %v`, client.HostURL, err)
	}
}
