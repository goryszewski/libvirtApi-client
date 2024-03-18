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

type LB struct {
	Id      int
	Ip      string
	service string
}

type portlb struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	NodePort int    `json:"nodeport"`
}

type bindServiceLB struct {
	Ports []portlb `json:"ports"`
}
