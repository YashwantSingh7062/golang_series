package models

type Config struct {
	Address string
}

func NewConfig(address string) *Config {
	cfg := &Config{
		Address: address,
	}

	return cfg
}
