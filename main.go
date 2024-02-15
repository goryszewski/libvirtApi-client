package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

	client, err := libvirtApiClient.NewClient(conf)

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
		log.Fatalf("[Error][main][004] BindLB return: ", err)
	}

	err = client.UnBindLB(service)
	if err != nil {
		log.Fatalf("[Error][main][005] UnBindLB return: ", err)
	}

	err = client.UpdateBind(service, nodes)
	if err != nil {
		log.Fatalf("[Error][main][006] UpdateBind return: ", err)
	}

}
