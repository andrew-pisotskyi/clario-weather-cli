package services

import (
	"errors"
	"sync"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/domain"
)

type WeatherService struct {
	Providers []domain.WeatherProvider
}

func NewWeatherService(providers []domain.WeatherProvider) *WeatherService {
	return &WeatherService{Providers: providers}
}

func (ws *WeatherService) GetFastestWeather(country, city string) (domain.Weather, error) {
	var wg sync.WaitGroup
	result := make(chan domain.Weather, len(ws.Providers))
	errorChan := make(chan error, len(ws.Providers))

	for _, provider := range ws.Providers {
		wg.Add(1)
		go func(p domain.WeatherProvider) {
			defer wg.Done()
			weather, err := p.GetWeather(country, city)
			if err != nil {
				errorChan <- err
				return
			}
			result <- weather
		}(provider)
	}

	go func() {
		wg.Wait()
		close(result)
		close(errorChan)
	}()

	var failedProviders int
	for {
		select {
		case weather := <-result:
			return weather, nil
		case err := <-errorChan:
			if err != nil {
				failedProviders++
			}
			if failedProviders == len(ws.Providers) {
				return domain.Weather{}, errors.New("all providers failed")
			}
		}
	}
}
