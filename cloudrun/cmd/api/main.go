package main

import (
	"github.com/mathcale/goexpert-course/cloudrun/config"
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/dependencies"
)

func main() {
	configs, configsErr := config.LoadConfig(".")
	if configsErr != nil {
		panic(configsErr)
	}

	deps := dependencies.Build(configs)
	deps.WebServer.Start()
}
