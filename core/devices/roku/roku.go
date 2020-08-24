package roku

import (
	"fmt"
	"github.com/rebelit/gome-core/common/httpRequest"
)

type Device interface {
	GetState() (state bool, error error)
	ControlPowerState(state bool) error
}

type Client struct {
	Name    string
	Id      string
	Address string
	Port    string
}

func sendKeyPress(url string) error {
	resp, err := httpRequest.Post(url, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ControlPowerState(state bool) error {
	var control = ""
	if state {
		control = "PowerOn"
	} else {
		control = "PowerOff"
	}

	url := fmt.Sprintf("http://%s:%s/%s/%s", c.Address, c.Port, "keypress", control)
	if err := sendKeyPress(url); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetState() (state bool, error error) {
	url := fmt.Sprintf("http://%s:%s/", c.Address, c.Port)

	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, fmt.Errorf("non-200 response %d", response.StatusCode)
	}

	return true, nil
}
