package config

type Config struct {
	Port string
}

func NewConfig() *Config {
	cfg := &Config{
		Port: ":3000",
	}

	return cfg
}
