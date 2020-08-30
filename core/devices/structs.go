package devices

type Devices struct {
	Device []Device `json:"device"`
}

type Device struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type DeviceType struct {
	Name string
}
