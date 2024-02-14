package main

import (
	"fmt"

	"github.com/goryszewski/libvirtApi-client/libvirtApiClient"
)

func main() {
	url := "http://127.0.0.1:8050"
	user := "test"
	pass := "test"
	client, err := libvirtApiClient.NewClient(&url, &user, &pass)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	client.SignIn()
}
