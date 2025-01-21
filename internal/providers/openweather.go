package providers

import (
	"encoding/json"
	"fmt"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/utils"
)

// OpenWeather is an implementation of the Open Weather Map API
type OpenWeather struct {
	apiKey     string
	httpClient *utils.HTTPClient
}

// NewOpenWeather creates an instance of OpenWeather
func NewOpenWeather(apiKey string, httpClient *utils.HTTPClient) *OpenWeather {
	return &OpenWeather{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// API response structure
type openWeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// GetWeather requests weather data from the Open Weather Map API
func (ow *OpenWeather) GetWeather(country, city string) (domain.Weather, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=%s",
		city, country, ow.apiKey,
	)
	response, err := ow.httpClient.Get(url)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to fetch data from OpenWeather API: %v", err)
	}

	var weatherData openWeatherResponse
	err = json.Unmarshal(response, &weatherData)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to parse OpenWeather API response: %v", err)
	}

	return domain.Weather{
		Temperature: weatherData.Main.Temp,
		Condition:   weatherData.Weather[0].Description,
		Provider:    "OpenWeatherMap",
	}, nil
}
