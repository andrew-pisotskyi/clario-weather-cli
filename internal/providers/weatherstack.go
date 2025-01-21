package providers

import (
	"encoding/json"
	"fmt"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/utils"
)

// WeatherStack is an implementation of the WeatherStack API
type WeatherStack struct {
	apiKey     string
	httpClient *utils.HTTPClient
}

// API response structure
type weatherStackResponse struct {
	Current struct {
		Temp      float64  `json:"temperature"`
		Condition []string `json:"weather_descriptions"`
		WindSpeed float64  `json:"wind_speed"`
	} `json:"current"`
}

// NewWeatherStack creates an instance of WeatherStack
func NewWeatherStack(apiKey string, httpClient *utils.HTTPClient) *WeatherStack {
	return &WeatherStack{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// GetWeather requests weather data from the WeatherStack API
func (ws *WeatherStack) GetWeather(country, city string) (domain.Weather, error) {
	url := fmt.Sprintf(
		"http://api.weatherstack.com/current?access_key=%s&query=%s,%s&units=m",
		ws.apiKey, city, country,
	)
	response, err := ws.httpClient.Get(url)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to fetch data from WeatherStack API: %v", err)
	}

	var weatherData weatherStackResponse
	err = json.Unmarshal(response, &weatherData)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to parse WeatherStack API response: %v", err)
	}

	return domain.Weather{
		Temperature: weatherData.Current.Temp,
		Condition:   weatherData.Current.Condition[0],
		WindSpeed:   weatherData.Current.WindSpeed,
		Provider:    "WeatherStack",
	}, nil
}
