package config

import "os"

type Config struct {
	DBPath string
}

func LoadConfig() *Config {
	return &Config{
		DBPath: getEnv("DB_PATH", "users.db"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
