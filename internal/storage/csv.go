package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
)

func SaveRatesToCSV(rates []api.ExchangeRate, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Date", "Currency", "Buy", "Sell"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing CSV header: %v", err)
	}

	// Write data
	date := time.Now().Format("2006-01-02")
	for _, rate := range rates {
		row := []string{
			date,
			rate.Currency,
			rate.Buy,
			rate.Sale,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing CSV row: %v", err)
		}
	}

	return nil
}

func SaveMonthlyRatesToCSV(rates map[string][]api.ExchangeRate, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Date", "Currency", "Buy", "Sell"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing CSV header: %v", err)
	}

	// Write data
	for date, dailyRates := range rates {
		for _, rate := range dailyRates {
			row := []string{
				date,
				rate.Currency,
				rate.Buy,
				rate.Sale,
			}
			if err := writer.Write(row); err != nil {
				return fmt.Errorf("error writing CSV row: %v", err)
			}
		}
	}

	return nil
}
