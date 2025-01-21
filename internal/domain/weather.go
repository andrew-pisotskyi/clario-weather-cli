package domain

// Weather - contains weather results
type Weather struct {
	Temperature float64
	Condition   string
	Provider    string
}

// WeatherProvider is an interface for weather providers
type WeatherProvider interface {
	GetWeather(country, city string) (Weather, error)
}
