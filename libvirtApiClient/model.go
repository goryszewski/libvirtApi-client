package libvirtApiClient

type Worker struct { // TODO refactor / use node
	Name     string `json:"name"`
	Type     string `json:"type"`
	Internal string `json:"internal"`
	External string `json:"external"`
}

type Port_Service struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	NodePort int    `json:"nodeport"`
}
type Node struct {
	Name     string `json:"name"`
	Internal string `json:"internal"`
	External string `json:"external"`
}

type LoadBalancer struct {
	Ports     []Port_Service `json:"ports"`
	Nodes     []Node         `json:"nodes"`
	Namespace string         `json:"namespace"`
	Name      string         `json:"name"`
	Ip        string         `json:"ip,omitempty"`
	//	Id        struct {
	//		OID string `json:"skipempty,omitempty"`
	//	} `json:"skipempty,omitempty"`
}

type Disk struct {
	ID    int `json:"id"`
	VM_ID int `json:"vm_id"`
	Size  int `json:"size"`
}

type BindDiskResponse struct {
	ID      string `json:"id"`
	Path    string `json:"path"`
	Target  string `json:"target"`
	Address string `json:"address"`
}
