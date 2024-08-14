package main

import (
	"log"
	"os"

	"github.com/Nikolay200669/CurrencyRate/internal/config"
	"github.com/Nikolay200669/CurrencyRate/internal/tray"
	"github.com/Nikolay200669/CurrencyRate/pkg/autostart"
	"github.com/getlantern/systray"
)

func main() {
	// Parse command-line arguments for autostart management
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--enable-autostart":
			err := autostart.Enable()
			if err != nil {
				log.Fatalf("Failed to enable autostart: %v", err)
			}
			os.Exit(0)
		case "--disable-autostart":
			err := autostart.Disable()
			if err != nil {
				log.Fatalf("Failed to disable autostart: %v", err)
			}
			os.Exit(0)
		}
	}

	// Load configuration
	cfg, err := config.LoadConfig("configs/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Start the system tray application
	systray.Run(onReady(cfg), onExit)
}

func onReady(cfg *config.Config) func() {
	return func() {
		tray.Initialize(cfg)
	}
}

func onExit() {
	// Perform cleanup tasks if necessary
}
