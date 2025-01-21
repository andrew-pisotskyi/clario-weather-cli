package config

import (
	"log"
	"os"
)

// Config contains the configuration of the application package
type Config struct {
	OpenWeatherAPIKey  string
	WeatherAPIKey      string
	WeatherStackAPIKey string
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{
		OpenWeatherAPIKey:  getEnv("OPEN_WEATHER_API_KEY", "86a199875f05dbcb2e6966fbb2aaa104"),
		WeatherAPIKey:      getEnv("WEATHER_API_KEY", "cdc49270f649408980e152042252001"),
		WeatherStackAPIKey: getEnv("WEATHER_STACK_API_KEY", "12caa2d7377313095708134f261e933f"),
	}
}

// getEnv returns an environment variable or a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s is required", key)
		}
		return defaultValue
	}
	return value
}
