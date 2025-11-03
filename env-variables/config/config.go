package config

type Config struct {
	Port string `env:"PORT" env-default:"3000"`
	Host string `env:"HOST" env-default:"localhost"`
}

func (c *Config) Address() string {
	return c.Host + ":" + c.Port
}
