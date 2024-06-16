package libvirtApiClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetDisks() error {
	request, err := http.NewRequest("GET", fmt.Sprintf("%v/api/v2/hdd", c.HostURL), nil)

	if err != nil {
		log.Fatal("Error")
	}

	body, _, err := c.doRequest(request)

	var disk []Disk

	err = json.Unmarshal(body, &disk)

	fmt.Printf("[%v]", disk)

	return nil
}
