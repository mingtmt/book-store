package configs

type Config struct {
	ServerAddress string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: "localhost:8080",
	}
}
