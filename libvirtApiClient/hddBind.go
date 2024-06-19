package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) BindDisk(disk_id int, node_id string) (*BindDiskResponse, error) {

	request, err := http.NewRequest("PUT", fmt.Sprintf("%v/api/v2/node/%v/hdd/%v", c.HostURL, node_id, disk_id), nil)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	body, _, err := c.doRequest(request)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	var diskinfo BindDiskResponse
	err = json.Unmarshal(body, &diskinfo)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return &diskinfo, nil
}
func (c *Client) UnBindDisk(disk_id int, node_id string) error {

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/v2/node/%v/hdd/%v", c.HostURL, node_id, disk_id), nil)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	_, _, err = c.doRequest(request)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	return nil
}
