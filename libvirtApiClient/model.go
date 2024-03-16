package libvirtApiClient

import v1 "k8s.io/api/core/v1"

type Ip struct {
	Private string
	Public  string
}

type Worker struct {
	Name string
	IP   Ip
	Type string
}
type Service struct {
	name      string
	namespace string
	port	  int
}

type LB struct {
	Id      int
	Ip      string
	service string
}

type bindServiceLB struct {
	service *v1.Service
	workers []*v1.Node
}
