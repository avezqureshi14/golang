package config

type Config struct {
	ServiceName string
	Env         string
	Version     string
}

func Load() Config {
	return Config{
		ServiceName: "go-backend",
		Env:         "dev",
		Version:     "1.0.0",
	}
}
