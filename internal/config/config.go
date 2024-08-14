package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	UpdateInterval  int      `json:"update_interval"`
	SaveFormats     string   `json:"save_formats"`
	Currencies      []string `json:"currencies"`
	CurrentRatesURL string   `json:"current_rates_url"`
	MonthlyRatesURL string   `json:"monthly_rates_url"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
