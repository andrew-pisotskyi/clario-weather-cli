package main

import (
	"fmt"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/config"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/providers"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/services"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/utils"
)

func main() {
	cfg := config.NewConfig()

	fmt.Println("Hi! Do you want to know the weather in your city? Let's go!")

	fmt.Print("Print your country code: ")
	var country string
	fmt.Scanln(&country)

	fmt.Print("Print your city: ")
	var city string
	fmt.Scanln(&city)

	httpClient := utils.NewHTTPClient(cfg.HTTPTimeout)
	providers := []domain.WeatherProvider{
		providers.NewOpenWeather(cfg.ProvidersKeys.OpenWeatherAPIKey, httpClient),
		providers.NewWeatherApi(cfg.ProvidersKeys.WeatherAPIKey, httpClient),
		providers.NewWeatherStack(cfg.ProvidersKeys.WeatherStackAPIKey, httpClient),
	}

	weatherService := services.NewWeatherService(providers)
	weather, err := weatherService.GetFastestWeather(country, city)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}

	fmt.Printf("Weather in %s, %s:\n", city, country)
	fmt.Printf("Temperature: %.2f°C\n", weather.Temperature)
	fmt.Printf("Condition: %s\n", weather.Condition)
	fmt.Printf("Provider: %s\n", weather.Provider)
}
