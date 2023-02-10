package config

func NewConfig() *Config {
	cfg := &Config{
		Http: HttpConfig{
			Port: ":3000",
		},

		Controller: ControllerConfig{
			StaticDir:   "./static",
			SpaViewRoot: "/spa_app/",
		},
	}

	return cfg
}
