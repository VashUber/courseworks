package config

type HttpConfig struct {
	Port string
}

type ControllerConfig struct {
	StaticPath string
}

type Config struct {
	Http       HttpConfig
	Controller ControllerConfig
}
