package logger

type Config struct {
	Env string
}

func NewConfig() Config {
	return Config{}
}
