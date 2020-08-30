package rpiot

import (
	"fmt"
	"github.com/rebelit/gome-core/common/httpRequest"
)

//ToDo:  complete this package... need a rpIoT go-client from swagger 1st.
type Device interface {
	GetState() (state bool, error error)
	Reboot() error
}

type Client struct {
	Address string
	Port    string
	Meta    Meta
}

type Meta struct {
	Gpio   Pin
	Action string
}

type Pin struct {
	Name   string
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
