package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var config Config
var once sync.Once

func loadConfig() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Warning : No .env file found")
	}

	Version := os.Getenv("VERSION")
	ServiceName := os.Getenv("SERVICE_NAME")
	HttpPort := os.Getenv("HTTP_PORT")

	if Version == "" {
		fmt.Println("Version is required")
		os.Exit(1)

	}

	if ServiceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	if HttpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(HttpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	config = Config{
		Version:     Version,
		ServiceName: ServiceName,
		HttpPort:    int(port),
	}

}

func GetConfig() Config {
	once.Do(func() {
		loadConfig()
	})
	return config
}
