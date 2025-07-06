package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Server struct {
	Address string
}

type Config struct {
	Server Server
}

func Load() *Config {
	return &Config{
		Server: Server{
			Address: fmt.Sprintf(":%d", getInt("SERVER_PORT", 9000)),
		},
	}
}

func PostgresConnectURL() string {
	url := getString("POSTGRES_URL", "")
	if url == "" {
		log.Fatalln("POSTGRES_URL отсутствует")
	}
	return url
}

func getString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {

		return defaultValue
	}
	return v
}

func getInt(key string, defaultValue int) int {
	v := os.Getenv(key)

	n, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return n
}
