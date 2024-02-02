package main

import (
	"assessment-go-source-code-muhammad-aditya/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	// ctx, cancel := context.WithCancel(context.Background())
}
