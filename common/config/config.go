package config

type Conf struct {
	StatAddr     string
	SlackWebhook string //https://hooks.slack.com/services/<ID>
	DbPath       string
	Name         string
	AuthToken    string
}

var App *Conf

//ToDo: load these from flag.  Static for development.
//ToDo: flags and defaults
func Runtime() {
	c := Conf{
		StatAddr:     "",
		SlackWebhook: "hooks.slack.com/guid/or/something",
		//DbPath:       "/usr/local/gome-core/db",
		DbPath:    "mocks", //local testing in development
		Name:      "gome-core",
		AuthToken: "test",
	}

	App = &c
	return
}
