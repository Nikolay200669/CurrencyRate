package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nikolay200669/CurrencyRate/internal/config"
)

type ExchangeRate struct {
	Currency     string `json:"ccy"`
	BaseCurrency string `json:"base_ccy"`
	Buy          string `json:"buy"`
	Sale         string `json:"sale"`
}

func GetCurrentRates(cfg *config.Config, currency []string) ([]ExchangeRate, error) {
	resp, err := http.Get(cfg.CurrentRatesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rates []ExchangeRate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, err
	}

	return filterRates(rates, currency), nil
}

func GetMonthlyRates(cfg *config.Config) (map[string][]ExchangeRate, error) {
	monthlyRates := make(map[string][]ExchangeRate)
	today := time.Now()

	for i := 0; i < 30; i++ {
		date := today.AddDate(0, 0, -i)
		url := fmt.Sprintf("%s%s", cfg.MonthlyRatesURL, date.Format("02.01.2006"))

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var ratesResponse struct {
			Date         string         `json:"date"`
			ExchangeRate []ExchangeRate `json:"exchangeRate"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&ratesResponse); err != nil {
			return nil, err
		}

		monthlyRates[ratesResponse.Date] = filterRates(ratesResponse.ExchangeRate, cfg.Currencies)
	}

	return monthlyRates, nil
}

func filterRates(rates []ExchangeRate, currencies []string) []ExchangeRate {
	var filteredRates []ExchangeRate
	for _, rate := range rates {
		for _, currency := range currencies {
			if rate.Currency == currency {
				filteredRates = append(filteredRates, rate)
				break
			}
		}
	}
	return filteredRates
}
