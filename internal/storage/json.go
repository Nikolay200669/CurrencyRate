package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
)

type RateData struct {
	Date  string             `json:"date"`
	Rates []api.ExchangeRate `json:"rates"`
}

func SaveRatesToJSON(rates []api.ExchangeRate, filename string) error {
	data := RateData{
		Date:  time.Now().Format("2006-01-02"),
		Rates: rates,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON file: %v", err)
	}

	return nil
}

func SaveMonthlyRatesToJSON(rates map[string][]api.ExchangeRate, filename string) error {
	var monthlyData []RateData

	for date, dailyRates := range rates {
		data := RateData{
			Date:  date,
			Rates: dailyRates,
		}
		monthlyData = append(monthlyData, data)
	}

	jsonData, err := json.MarshalIndent(monthlyData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON file: %v", err)
	}

	return nil
}
