package config

type HttpConfig struct {
	Port string
}

type ControllerConfig struct {
	StaticDir   string
	SpaViewRoot string
}

type Config struct {
	Http       HttpConfig
	Controller ControllerConfig
}
