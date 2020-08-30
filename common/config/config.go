package config

import (
	"flag"
	"log"
)

type Conf struct {
	StatAddr     string //127.0.0.1:8125
	SlackWebhook string //https://hooks.slack.com/services/<ID>
	DbPath       string
	Name         string
	AuthToken    string
	ListenPort   string
}

var App *Conf

func Runtime() {
	c := &Conf{}
	log.Printf("empty %+v", c)
	configDefaults(c)
	log.Printf("defaults %+v", c)
	configFlags(c)
	log.Printf("flags %+v", c)

	App = c
	return
}

func configDefaults(c *Conf) {
	c.Name = "gome-core"
	c.StatAddr = ""
	c.DbPath = "/usr/local/gome-core/db"
	c.AuthToken = "changeMePlease"
	c.ListenPort = "6660"
	return
}

func configFlags(c *Conf) {
	flag.StringVar(&c.StatAddr, "statsd", c.StatAddr, "statsd address")
	flag.StringVar(&c.Name, "name", c.Name, "application name")
	flag.StringVar(&c.DbPath, "dbPath", c.DbPath, "path to local database")
	flag.StringVar(&c.SlackWebhook, "slackWebhook", c.SlackWebhook, "slack webhook url")
	flag.StringVar(&c.AuthToken, "authToken", c.AuthToken, "app authentication token")
	flag.StringVar(&c.ListenPort, "port", c.ListenPort, "http listener http port")
	flag.Parse()
	return
}
