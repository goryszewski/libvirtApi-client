package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetDisks() (*[]Disk, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/v2/hdd", c.HostURL), nil)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	body, _, err := c.doRequest(request)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	var disks []Disk

	err = json.Unmarshal(body, &disks)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return &disks, nil
}

func (c *Client) GetDisk(id int) (*Disk, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/v2/hdd/%v", c.HostURL, id), nil)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	body, _, err := c.doRequest(request)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	var disk Disk

	err = json.Unmarshal(body, &disk)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return &disk, nil
}

func (c *Client) DeleteDisk(id int) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%v/api/v2/hdd/%v", c.HostURL, id), nil)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	_, _, err = c.doRequest(request)

	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	return nil
}

func (c *Client) CreateDisk(size int) (*Disk, error) {
	size_string := strconv.Itoa(size)
	payload, err := json.Marshal(map[string]string{"size": size_string, "path": "/var/lib/libvirt/images"})

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v2/hdd", c.HostURL), strings.NewReader(string(payload)))

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	body, _, err := c.doRequest(request)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	var disk Disk

	err = json.Unmarshal(body, &disk)

	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return &disk, nil
}
