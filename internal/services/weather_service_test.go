package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
	"github.com/andrew-pisotskyi/clario-weather-cli/internal/services"
)

type MockWeatherProvider struct {
	Weather domain.Weather
	Error   error
	Timeout time.Duration
}

func (m *MockWeatherProvider) GetWeather(country, city string) (domain.Weather, error) {
	if m.Timeout > 0 {
		time.Sleep(m.Timeout)
	}
	return m.Weather, m.Error
}

func TestGetFastestWeather_Success(t *testing.T) {
	mockProvider1 := &MockWeatherProvider{
		Weather: domain.Weather{Temperature: 25, Condition: "Sunny", Provider: "Provider1"},
		Error:   nil,
		Timeout: 0 * time.Millisecond,
	}

	mockProvider2 := &MockWeatherProvider{
		Weather: domain.Weather{Temperature: 30, Condition: "Sunny", Provider: "Provider2"},
		Error:   nil,
		Timeout: 500 * time.Millisecond,
	}

	ws := services.NewWeatherService([]domain.WeatherProvider{mockProvider1, mockProvider2})

	weather, err := ws.GetFastestWeather("UA", "Kyiv")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if weather.Temperature != 25 {
		t.Fatalf("expected temperature 25, got %v", weather.Temperature)
	}
}

func TestGetFastestWeather_AllFail(t *testing.T) {
	mockProvider1 := &MockWeatherProvider{
		Weather: domain.Weather{},
		Error:   errors.New("network error"),
		Timeout: 0 * time.Millisecond,
	}

	mockProvider2 := &MockWeatherProvider{
		Weather: domain.Weather{},
		Error:   errors.New("city not found"),
		Timeout: 500 * time.Millisecond,
	}

	ws := services.NewWeatherService([]domain.WeatherProvider{mockProvider1, mockProvider2})

	_, err := ws.GetFastestWeather("Ukraine", "Kyiv")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, domain.ErrAllProvidesFailed) {
		t.Fatalf("expected error: %v, got %v", domain.ErrAllProvidesFailed, err)
	}
}
