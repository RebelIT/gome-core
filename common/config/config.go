package config

type Conf struct {
	StatAddr string
}

var App *Conf

func Runtime() {
	c := Conf{
		StatAddr: "",
	}

	App = &c
	return
}
