package config

import "flag"

type Conf struct {
	StatAddr     string
	SlackWebhook string //https://hooks.slack.com/services/<ID>
	DbPath       string
	Name         string
	AuthToken    string
}

var App *Conf

func Runtime() {
	c := &Conf{}

	configDefaults(c)
	configFlags(c)

	App = c
	return
}

func configDefaults(c *Conf) {
	c.Name = "gome-core"
	c.StatAddr = "127.0.0.1:8125"
	c.DbPath = "/usr/local/gome-core/db"
	c.AuthToken = "changeMePlease"
}

func configFlags(c *Conf) {
	flag.StringVar(&c.StatAddr, "statsd", "", "statsd address")
	flag.StringVar(&c.Name, "name", "", "application name")
	flag.StringVar(&c.DbPath, "dbPath", "", "path to local database")
	flag.StringVar(&c.SlackWebhook, "slackWebhook", "", "slack webhook url")
	flag.StringVar(&c.AuthToken, "authToken", "", "app authentication token")
	flag.Parse()
}
