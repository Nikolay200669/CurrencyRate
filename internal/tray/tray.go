package tray

import (
	"fmt"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
	"github.com/Nikolay200669/CurrencyRate/internal/config"
	"github.com/Nikolay200669/CurrencyRate/internal/currency"
	"github.com/Nikolay200669/CurrencyRate/internal/storage"

	"github.com/getlantern/systray"
)

func Initialize(cfg *config.Config) {
	systray.SetIcon(getIcon())
	systray.SetTitle("Currency App")
	systray.SetTooltip("Currency Exchange Rates")

	mRefresh := systray.AddMenuItem("Refresh Rates", "Get current exchange rates")
	mMonthly := systray.AddMenuItem("Monthly Rates", "Get exchange rates for the last month")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for {
			select {
			case <-mRefresh.ClickedCh:
				refreshRates(cfg)
			case <-mMonthly.ClickedCh:
				getMonthlyRates(cfg)
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func refreshRates(cfg *config.Config) {
	rates, err := api.GetCurrentRates(cfg.Currencies)
	if err != nil {
		fmt.Println("Error fetching current rates:", err)
		return
	}

	err = storage.SaveRates(rates, cfg.SaveFormats)
	if err != nil {
		fmt.Println("Error saving rates:", err)
	}

	currency.UpdateTrayMenu(rates)
}

func getMonthlyRates(cfg *config.Config) {
	rates, err := api.GetMonthlyRates(cfg.Currencies)
	if err != nil {
		fmt.Println("Error fetching monthly rates:", err)
		return
	}

	err = storage.SaveMonthlyRates(rates, cfg.SaveFormats)
	if err != nil {
		fmt.Println("Error saving monthly rates:", err)
	}

	currency.DisplayMonthlyRates(rates)
}

func getIcon() []byte {
	// Implementation to load and return the icon bytes
	// This could be reading from a file in the assets directory
	return []byte{}
}
