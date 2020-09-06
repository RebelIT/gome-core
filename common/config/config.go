package config

import (
	"flag"
	"log"
	"os"
)

type Conf struct {
	StatAddr     string //127.0.0.1:8125
	SlackWebhook string //https://hooks.slack.com/services/<ID>
	DbPath       string
	Name         string
	AuthToken    string
	ListenPort   string
	GenerateSpec bool
	FullMemory   bool
}

var App *Conf

func Runtime() {
	log.Printf("INFO: loading runtime configuration")
	c := &Conf{}
	configDefaults(c)
	configEnvironment(c)
	configFlags(c)

	App = c
	return
}

func configDefaults(c *Conf) {
	c.Name = "gome-core"
	c.StatAddr = ""
	c.DbPath = "badgerDatabase"
	c.AuthToken = "changeMePlease"
	c.ListenPort = "6660"
	c.GenerateSpec = false
	c.FullMemory = false
	return
}

func configEnvironment(c *Conf) {
	name := os.Getenv("CORE_NAME")
	statsd := os.Getenv("CORE_STATSD")
	dbPath := os.Getenv("CORE_DBPATH")
	slackWebhook := os.Getenv("CORE_SLACK")
	authToken := os.Getenv("CORE_TOKEN")
	port := os.Getenv("CORE_PORT")
	fullMemory := os.Getenv("CORE_MEMORY")

	if name != "" {
		c.Name = name
	}
	if statsd != "" {
		c.StatAddr = statsd
	}
	if dbPath != "" {
		c.DbPath = dbPath
	}
	if slackWebhook != "" {
		c.SlackWebhook = slackWebhook
	}
	if authToken != "" {
		c.AuthToken = authToken
	}
	if port != "" {
		c.ListenPort = port
	}
	if fullMemory != "" {
		c.FullMemory = true
	}

	return
}

func configFlags(c *Conf) {
	flag.StringVar(&c.StatAddr, "statsd", c.StatAddr, "statsd address")
	flag.StringVar(&c.Name, "name", c.Name, "application name")
	flag.StringVar(&c.DbPath, "dbPath", c.DbPath, "path to local database")
	flag.StringVar(&c.SlackWebhook, "slackWebhook", c.SlackWebhook, "slack webhook url")
	flag.StringVar(&c.AuthToken, "authToken", c.AuthToken, "app authentication token")
	flag.StringVar(&c.ListenPort, "port", c.ListenPort, "http listener http port")
	flag.BoolVar(&c.GenerateSpec, "generateSpec", c.GenerateSpec, "print the http spec to console")
	flag.BoolVar(&c.FullMemory, "fullMemory", c.FullMemory, "run database with full memory cache mem > 1GB required")
	flag.Parse()
	return
}
