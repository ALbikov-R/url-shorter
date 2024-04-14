package storage

import "fmt"

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	SSLmode  string
}

func (c PostgresConfig) string() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBname, c.SSLmode)
}

type RedisConfig struct {
	URL  string
	Port int
}

func (c RedisConfig) string() string {
	return fmt.Sprintf("%s:%d", c.URL, c.Port)
}
