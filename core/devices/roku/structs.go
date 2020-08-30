package roku

import "encoding/xml"

//device info response
type DeviceInfo struct {
	XMLName            xml.Name `xml:"device-info"`
	Text               string   `xml:",chardata"`
	Udn                string   `xml:"udn"`
	SerialNumber       string   `xml:"serial-number"`
	DeviceID           string   `xml:"device-id"`
	IsTv               string   `xml:"is-tv"`
	IsStick            string   `xml:"is-stick"`
	NetworkName        string   `xml:"network-name"`
	UserDeviceName     string   `xml:"user-device-name"`
	UserDeviceLocation string   `xml:"user-device-location"`
	Uptime             string   `xml:"uptime"`
	PowerMode          string   `xml:"power-mode"`
}

//apps queries response
//installed App query
type Apps struct {
	XMLName xml.Name `xml:"apps"`
	Text    string   `xml:",chardata"`
	App     []App    `xml:"app"`
}

//active app response
type ActiveApp struct {
	XMLName xml.Name `xml:"active-app"`
	Text    string   `xml:",chardata"`
	App     App      `xml:"app"`
}

type App struct {
	Text    string `xml:",chardata"`
	ID      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
}

//handler responses
type RespPwr struct {
	Name  string `json:"name"`
	State bool   `json:"power_state"`
}

type JsonDeviceInfo struct {
	Udn                string `json:"udn"`
	SerialNumber       string `json:"serial_number"`
	DeviceID           string `json:"device_id"`
	IsTv               string `json:"is_tv"`
	IsStick            string `json:"is_stick"`
	NetworkName        string `json:"network_name"`
	UserDeviceName     string `json:"user_device_name"`
	UserDeviceLocation string `json:"user_device_location"`
	Uptime             string `json:"uptime"`
	PowerMode          string `json:"power_mode"`
}

type JsonApps struct {
	Apps []JsonApp `json:"app"`
}

type JsonActiveApp struct {
	App JsonApp `json:"app"`
}

type JsonApp struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Version string `json:"version"`
}
