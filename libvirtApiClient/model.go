package libvirtApiClient

type Ip struct {
	Private string
	Public  string
}

type Worker struct {
	Name string
	IP   Ip
	Type string
}

type LoadBalancer struct {
	Id int
	Ip string
}

type Port_Service struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	NodePort int    `json:"nodeport"`
}
type Node struct {
	Name       string `json:"name"`
	Public_ip  string `json:"public_ip"`
	Private_ip string `json:"private_ip"`
}

type ServiceLoadBalancer struct {
	Ports     []Port_Service `json:"ports"`
	Name      string         `json:"name"`
	Namespace string         `json:"namespace"`
	Nodes     []Node         `json:"nodes"`
}
type ServiceLoadBalancerResponse struct {
	ID any    `json:"_id"`
	Ip string `json:"ip"`
	*ServiceLoadBalancer
}
