package currency

import (
	"fmt"
	"strings"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
	"github.com/getlantern/systray"
)

var currencyMenuItems map[string]*systray.MenuItem

func InitializeCurrencyMenu(currencies []string) {
	currencyMenuItems = make(map[string]*systray.MenuItem)
	for _, currency := range currencies {
		menuItem := systray.AddMenuItem(currency, fmt.Sprintf("Show %s exchange rate", currency))
		currencyMenuItems[currency] = menuItem
	}
}

func UpdateTrayMenu(rates []api.ExchangeRate) {
	for _, rate := range rates {
		if menuItem, exists := currencyMenuItems[rate.Currency]; exists {
			menuItem.SetTitle(formatRate(rate))
		}
	}
}

func DisplayMonthlyRates(rates map[string][]api.ExchangeRate) {
	// This function could be expanded to display monthly rates in a more user-friendly way
	// For now, we'll just print them to the console
	for date, dailyRates := range rates {
		fmt.Printf("Date: %s\n", date)
		for _, rate := range dailyRates {
			fmt.Printf("  %s\n", formatRate(rate))
		}
		fmt.Println()
	}
}

func formatRate(rate api.ExchangeRate) string {
	return fmt.Sprintf("%s: Buy %.2f / Sell %.2f",
		rate.Currency,
		parseFloat(rate.Buy),
		parseFloat(rate.Sale))
}

func parseFloat(s string) float64 {
	var f float64
	fmt.Sscanf(strings.TrimSpace(s), "%f", &f)
	return f
}
