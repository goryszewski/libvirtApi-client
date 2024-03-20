package main

import (
	"encoding/json"
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

	port := libvirtApiClient.Port_Service{
		Name:     "test",
		Protocol: "TCP",
		Port:     1,
		NodePort: 2,
	}
	var ports []libvirtApiClient.Port_Service

	node1 := libvirtApiClient.Node{
		Name:       "test",
		Private_ip: "10.10.11.1",
		Public_ip:  "192.168.1.1",
	}

	var nodes []libvirtApiClient.Node

	nodes = append(nodes, node1)
	ports = append(ports, port)

	bind_payload := libvirtApiClient.ServiceLoadBalancer{Name: "nnn21", Namespace: "nnn", Ports: ports, Nodes: nodes}

	ip, err := client.CreateLoadBalancer(bind_payload)

	if err != nil {
		log.Fatalf("[Error][main][004] BindLB return: %v", err)
	}

	log.Printf("ip: [%v]", ip)
	node2 := libvirtApiClient.Node{
		Name:       "test2",
		Private_ip: "10.10.11.1",
		Public_ip:  "192.168.1.1",
	}
	nodes = append(nodes, node2)

	update_payload := libvirtApiClient.ServiceLoadBalancer{Name: "nnn21", Namespace: "nnn", Ports: ports, Nodes: nodes}

	err = client.UpdateLoadBalancer(update_payload)
	if err != nil {
		log.Fatalf("[Error][main][006] UpdateBind return: %v", err)
	}
	log.Printf("update error: [%v]", err)

	err = client.DeleteLoadBalancer(bind_payload)
	if err != nil {
		log.Fatalf("[Error][main][005] UnBindLB return: %v", err)
	}
	log.Printf("delete error: [%v]", err)

}
