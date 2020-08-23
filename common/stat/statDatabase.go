package stat

import "gopkg.in/alexcesaro/statsd.v2"

func Database(action string, state string) {
	measurement := "database"
	tags := statsd.Tags("action", action, "state", state)

	counter(measurement, tags)
	return
}
