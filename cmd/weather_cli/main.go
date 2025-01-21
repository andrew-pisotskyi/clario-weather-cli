package main

import (
	"fmt"

	"github.com/andrew-pisotskyi/clario-weather-cli/internal/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println("cfg ", cfg)
}
