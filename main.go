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
		Name:     "test",
		Internal: "10.10.11.1",
		External: "192.168.1.1",
	}

	var nodes []libvirtApiClient.Node

	nodes = append(nodes, node1)
	ports = append(ports, port)

	// Test not exist lb
	not_exist := libvirtApiClient.LoadBalancer{Name: "not_exist", Namespace: "not_exist", Ports: ports, Nodes: nodes}
	loadbalancer, exist, err := client.GetLoadBalancer(not_exist)
	log.Printf("[%+v] | [%v] | error:[%v]", loadbalancer, exist, err)
	if err != nil {
		log.Fatalf("[Error][main][004] GetLoadBalancer return: %v", err)
	}

	bind_payload := libvirtApiClient.LoadBalancer{Name: "nnn21", Namespace: "nnn", Ports: ports, Nodes: nodes}
	ip, err := client.CreateLoadBalancer(bind_payload)

	if err != nil {
		log.Fatalf("[Error][main][004] CreateLoadBalancer return: %v", err)
	}

	log.Printf("ip: [%v]", ip)

	loadbalancer, exist, err = client.GetLoadBalancer(bind_payload)
	if err != nil {
		log.Fatalf("[Error][main][004] GetLoadBalancer return: %v", err)
	}
	log.Printf("ip: [%v][%v]", loadbalancer.Ip, exist)
	node2 := libvirtApiClient.Node{
		Name:     "test2",
		Internal: "10.10.11.1",
		External: "192.168.1.1",
	}
	nodes = append(nodes, node2)

	update_payload := libvirtApiClient.LoadBalancer{Name: "nnn21", Namespace: "nnn", Ports: ports, Nodes: nodes}

	err = client.UpdateLoadBalancer(update_payload)
	if err != nil {
		log.Fatalf("[Error][main][006] UpdateBind return Error: %v", err)
	}
	log.Printf("update ok: [%v]\n", err)

	err = client.DeleteLoadBalancer(bind_payload)
	if err != nil {
		log.Fatalf("[Error][main][005] UnBindLB return Error: %v", err)
	}
	log.Printf("delete ok: [%v]\n", err)

	// BEGIN DISK

	disks, err := client.GetDisks()
	if err != nil {
		log.Fatalf("[Error][main][GetDisks]  return Error: %v", err)
	}
	log.Printf("GetDisks ok: [%+v]\n", disks)

	disk, err := client.CreateDisk(10)
	if err != nil {
		log.Fatalf("[Error][main][CreateDisk]  return Error: %v", err)
	}
	log.Printf("CreateDisk ok: [%+v]\n", disk)

	// BEGIN BIND

	err = client.BindDisk(disk.ID, "worker01.autok8s.xyz")
	if err != nil {
		log.Fatalf("[Error][main][BindDisk] return Error: %v", err)
	}
	log.Printf("BindDisk ok: [%+v]\n", err)

	err = client.UnBindDisk(disk.ID, "worker01.autok8s.xyz")
	if err != nil {
		log.Fatalf("[Error][main][UnBindDisk]  return Error: %v", err)
	}
	log.Printf("UnBindDisk ok: [%+v]\n", err)

	// END BIND
	disks, err = client.GetDisks()
	if err != nil {
		log.Fatalf("[Error][main][GetDisks]  return Error: %v", err)
	}
	log.Printf("GetDisks ok: [%+v]\n", disks)

	err = client.DeleteDisk(disk.ID)
	if err != nil {
		log.Fatalf("[Error][main][DeleteDisk]  return Error: %v", err)
	}
	log.Printf("DeleteDisk ok: [%+v]\n", err)

	disks, err = client.GetDisks()
	if err != nil {
		log.Fatalf("[Error][main][GetDisks]  return Error: %v", err)
	}
	log.Printf("GetDisks ok: [%+v]\n", disks)

	// END DISK
}
