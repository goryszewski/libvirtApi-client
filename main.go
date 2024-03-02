package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/goryszewski/libvirtApi-client/libvirtApiClient"
)

func main() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("[Error][main][000] ReadFile return: %v", err)

	}
	var conf libvirtApiClient.Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("[Error][main][001] Unmarshal return: %v", err)
	}

	requester := &http.Client{Timeout: 10 * time.Second}

	client, err := libvirtApiClient.NewClient(conf, requester)

	if err != nil {
		log.Fatalf("[Error][main][002] NewClient return: %v", err)
	}

	client.SignIn()

	lb, err := client.GetFreeLB()
	if err != nil {
		log.Fatalf("[Error][main][003] GetFreeLB return: %v", err)
	}

	fmt.Println(lb)

	nodes := "test"
	service := "test_some_service"

	err = client.BindLB(lb.Ip, service, nodes)

	if err != nil {
		log.Fatalf("[Error][main][004] BindLB return: %v", err)
	}

	err = client.UnBindLB(service)
	if err != nil {
		log.Fatalf("[Error][main][005] UnBindLB return: %v", err)
	}

	err = client.UpdateBind(service, nodes)
	if err != nil {
		log.Fatalf("[Error][main][006] UpdateBind return: %v", err)
	}

}
