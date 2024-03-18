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

type ServiceLoadBalancer struct {
	Ports     []Port_Service `json:"ports"`
	Name      string         `json:"name"`
	Namespace string         `json:"namespace"`
}
