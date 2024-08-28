package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := NewYAMLHandler(os.Stdout, &HandlerOptions{})

	logger := slog.New(handler)

	logger.Info("Hello, world!")
	logger.Info("whatup!")
	logger.Info("Diner my shiny shiny love")
	logger.Info("All im thinking of", "ability", true)
}
