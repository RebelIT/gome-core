package stat

import (
	"github.com/rebelit/gome-core/common/config"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
)

func disabled() bool {
	if config.App.StatAddr == "" {
		return true
	}
	return false
}

// Generic metrics types using influx line protocol
func counter(measurement string, tags statsd.Option) {
	if disabled() {
		return
	}

	addrOpt := statsd.Address(config.App.StatAddr)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	s, err := statsd.New(addrOpt, fmtOpt, tags)
	if err != nil {
		log.Printf("ERROR sending %s counter %s", measurement, err)
	}
	defer s.Close()

	s.Increment(measurement)
	return
}

func gauge(measurement string, tags statsd.Option, value int) {
	if disabled() {
		return
	}

	addrOpt := statsd.Address(config.App.StatAddr)
	fmtOpt := statsd.TagsFormat(statsd.InfluxDB)
	s, err := statsd.New(addrOpt, fmtOpt, tags)
	if err != nil {
		log.Printf("ERROR sending %s gauge %s", measurement, err)
	}
	defer s.Close()

	s.Gauge(measurement, value)
	return
}
