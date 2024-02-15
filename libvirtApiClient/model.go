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
