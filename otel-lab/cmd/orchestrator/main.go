package main

import (
	"log"

	"github.com/mathcale/goexpert-course/otel-lab/config"
	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/dependencies"
)

func main() {
	configs, configsErr := config.LoadConfig(".")
	if configsErr != nil {
		log.Fatal(configsErr)
	}

	deps := dependencies.ResolveOrchestratorServiceDependencies(configs)
	deps.WebServer.Start()
}
