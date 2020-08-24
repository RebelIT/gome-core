package rpiot

import (
	"fmt"
	"github.com/rebelit/gome-core/common/httpRequest"
)

type Device interface {
	GetState() (state bool, error error)
	Reboot() error
}

type Client struct {
	Name    string
	Id      string
	Address string
	Port    string
	Gpio    Pin
}

type Pin struct {
	Number string
	Action string
}

func send(url string) error {
	resp, err := httpRequest.Post(url, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) Reboot() error {
	uriPart := "api/power/reboot"
	url := fmt.Sprintf("http://%s:%s/%s", c.Address, c.Port, uriPart)

	if err := send(url); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetState() (state bool, error error) {
	uriPart := "api/alive"
	url := fmt.Sprintf("http://%s:%s/%s", c.Address, c.Port, uriPart)

	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, fmt.Errorf("non-200 response %d", response.StatusCode)
	}

	return true, nil
}
