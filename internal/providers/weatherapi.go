package providers

import (
	"encoding/json"
	"fmt"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/utils"
)

// WeatherApi is an implementation of the Weather API
type WeatherApi struct {
	apiKey     string
	httpClient *utils.HTTPClient
}

// NewWeatherApi creates an instance of WeatherApi
func NewWeatherApi(apiKey string, httpClient *utils.HTTPClient) *WeatherApi {
	return &WeatherApi{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// API response structure
type weatherApiResponse struct {
	Current struct {
		Temp      float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Wind float64 `json:"wind_kph"`
	} `json:"current"`
}

// GetWeather requests weather data from the Weather API
func (w *WeatherApi) GetWeather(country, city string) (domain.Weather, error) {
	url := fmt.Sprintf(
		"http://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s",
		w.apiKey, city, country,
	)
	response, err := w.httpClient.Get(url)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to fetch data from Weather API: %v", err)
	}

	var weatherData weatherApiResponse
	err = json.Unmarshal(response, &weatherData)
	if err != nil {
		return domain.Weather{}, fmt.Errorf("failed to parse Weather API response: %v", err)
	}

	return domain.Weather{
		Temperature: weatherData.Current.Temp,
		Condition:   weatherData.Current.Condition.Text,
		WindSpeed:   weatherData.Current.Wind,
		Provider:    "WeatherAPI",
	}, nil
}
