package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetApplicationPort() int {
	port, err := strconv.Atoi(getEnvironmentValue("APPLICATION_PORT"))
	if err != nil {
		log.Fatalf("APPLICATION_PORT environment variable is not a number.")
	}

	return port
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func getEnvironmentValue(key string) string {
	log.Printf("Getting environment variable %s", os.Getenv(key))
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}
