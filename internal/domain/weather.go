package domain

// Weather - contains weather results
type Weather struct {
	Temperature float64
	Condition   string
	WindSpeed   float64
	Provider    string
}

// WeatherProvider is an interface for weather providers
type WeatherProvider interface {
	GetWeather(country, city string) (Weather, error)
}
