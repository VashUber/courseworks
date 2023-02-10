package config

func NewConfig() *Config {
	cfg := &Config{
		Http: HttpConfig{
			Port: ":4000",
		},

		Controller: ControllerConfig{
			StaticPath: "./static/index.html",
		},
	}

	return cfg
}
