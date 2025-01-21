package domain

type Weather struct {
	Temperature float64
	WindSpeed   float64
	Provider    string
}

type WeatherProvider interface {
	GetWeather(country, city string) (Weather, error)
}
