package config

type Conf struct {
	StatAddr     string
	SlackWebhook string //https://hooks.slack.com/services/<ID>
}

var App *Conf

func Runtime() {
	c := Conf{
		StatAddr:     "",
		SlackWebhook: "",
	}

	App = &c
	return
}
