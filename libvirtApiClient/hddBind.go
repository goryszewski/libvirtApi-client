package libvirtApiClient

import (
	"fmt"
	"net/http"
)

func (c *Client) BindDisk(disk_id int, node_id string) error {

	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/v2/hdd/%v/vm/%v", c.HostURL, disk_id, node_id), nil)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	_, _, err = c.doRequest(request)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	return nil
}
func (c *Client) UnBindDisk(disk_id int, node_id string) error {

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/v2/hdd/%v/vm/%v", c.HostURL, disk_id, node_id), nil)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	_, _, err = c.doRequest(request)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	return nil
}
