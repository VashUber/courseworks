package config

type Config struct {
	Port string
}

func NewConfig() *Config {
	cfg := &Config{
		Port: ":4000",
	}

	return cfg
}
