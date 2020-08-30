package roku

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rebelit/gome-core/common/config"
	"github.com/rebelit/gome-core/common/httpRequest"
	"github.com/rebelit/gome-core/database"
	"io/ioutil"
)

type Actions interface {
	getInfo() (deviceInfo DeviceInfo, error error) //get full device info
	getPowerState() (state bool, error error)      //get power on or off true|false
	controlPowerState(state bool) error            //set power on or off true|false
	getOnlineState() (state bool, error error)     //get online state by active ip address status code 200
	getApps() (apps Apps, error error)             //get all installed channels (apps)
	getActiveApp() (app App, error error)          //get currently launched channel (app)
	launchApp(appId string) error                  //launch an app by id
	keyPress(key string) error
}

type Client struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
}

const DBNAME = "roku"

var keys = []string{
	"home",
	"rev",
	"fwd",
	"play",
	"select",
	"left",
	"right",
	"down",
	"up",
	"back",
	"instantReplay",
	"info",
	"backspace",
	"search",
	"enter",
	"volumeDown",
	"volumeMute",
	"volumeUp",
	"powerOff",
}

//Interface Implements
func (c *Client) getInfo() (deviceInfo DeviceInfo, error error) {
	url := fmt.Sprintf("http://%s:%s/query/device-info", c.Address, c.Port)
	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return deviceInfo, err
	}
	if response.StatusCode != 200 {
		return deviceInfo, fmt.Errorf("non-200 response %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err := xml.Unmarshal(body, &deviceInfo); err != nil {
		return deviceInfo, err
	}

	return deviceInfo, nil
}

func (c *Client) getPowerState() (state bool, error error) {
	state = false
	deviceInfo := DeviceInfo{}

	url := fmt.Sprintf("http://%s:%s/query/device-info", c.Address, c.Port)
	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return state, err
	}
	if response.StatusCode != 200 {
		return state, fmt.Errorf("non-200 response %d", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err := xml.Unmarshal(body, &deviceInfo); err != nil {
		return state, err
	}

	if deviceInfo.PowerMode == "PowerOn" {
		state = true
	}

	return state, nil
}

func (c *Client) controlPowerState(state bool) error {
	var control = ""
	if state {
		control = "PowerOn"
	} else {
		control = "PowerOff"
	}

	url := fmt.Sprintf("http://%s:%s/keypress/%s", c.Address, c.Port, control)
	if err := sendKeyPress(url); err != nil {
		return err
	}

	return nil
}

func (c *Client) getOnlineState() (state bool, error error) {
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

func (c *Client) getApps() (apps Apps, error error) {
	url := fmt.Sprintf("http://%s:%s/query/apps", c.Address, c.Port)

	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return apps, err
	}
	if response.StatusCode != 200 {
		return apps, fmt.Errorf("non-200 response %d", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err := xml.Unmarshal(body, &apps); err != nil {
		return apps, err
	}

	return apps, nil
}

func (c *Client) getActiveApp() (app ActiveApp, error error) {
	url := fmt.Sprintf("http://%s:%s/query/active-app", c.Address, c.Port)

	response, err := httpRequest.Get(url, nil)
	if err != nil {
		return app, err
	}
	if response.StatusCode != 200 {
		return app, fmt.Errorf("non-200 response %d", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err := xml.Unmarshal(body, &app); err != nil {
		return app, err
	}

	return app, nil
}

func (c *Client) launchApp(appId string) error {
	url := fmt.Sprintf("http://%s:%s/launch/%s", c.Address, c.Port, appId)

	response, err := httpRequest.Post(url, nil, nil)
	if err != nil {
		return err
	}
	if response.StatusCode == 200 || response.StatusCode == 204 {
		return nil
	}

	return fmt.Errorf("non-200 return code %d", response.StatusCode)

}

func (c *Client) keyPress(key string) error {
	url := fmt.Sprintf("http://%s:%s/%s/%s", c.Address, c.Port, "keypress", key)
	if err := sendKeyPress(url); err != nil {
		return err
	}

	return nil
}

//private functions
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

func unmarshalClient(data []byte) (client Client) {
	err := json.Unmarshal(data, &client)
	if err != nil {
		return client
	}

	return client
}

//Public functions
func InitializeDb() error {
	db, err := database.Open(fmt.Sprintf("%s/%s", config.App.DbPath, DBNAME))
	if err != nil {
		return err
	}

	_, err = db.GetAllKeys()
	if err != nil {
		return err
	}

	return nil
}

func GetDeviceFromDb(name string) (client Client, error error) {
	db, err := database.Open(fmt.Sprintf("%s/%s", config.App.DbPath, DBNAME))
	if err != nil {
		return client, err
	}

	data, err := db.Get(name)
	if err != nil {
		return client, err
	}

	return unmarshalClient(data), nil
}

func GetAllDevicesFromDb() (clients []Client, error error) {
	db, err := database.Open(fmt.Sprintf("%s/%s", config.App.DbPath, DBNAME))
	if err != nil {
		return clients, err
	}

	keys, err := db.GetAllKeys()
	if err != nil {
		return clients, err
	}

	for _, key := range keys {
		client, err := GetDeviceFromDb(key)
		if err != nil {
			continue
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func LoadDevice(data []byte) error {
	client := unmarshalClient(data)

	db, err := database.Open(fmt.Sprintf("%s/%s", config.App.DbPath, DBNAME))
	if err != nil {
		return err
	}

	value, _ := json.Marshal(client)
	err = db.Set(client.Name, value)
	if err != nil {
		return err
	}

	return nil
}
