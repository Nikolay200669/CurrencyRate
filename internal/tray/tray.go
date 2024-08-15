package tray

import (
	"time"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
	"github.com/Nikolay200669/CurrencyRate/internal/config"
	"github.com/Nikolay200669/CurrencyRate/internal/currency"
	"github.com/Nikolay200669/CurrencyRate/internal/storage"
	"github.com/Nikolay200669/CurrencyRate/internal/utils"

	"github.com/getlantern/systray"
)

var (
	cfg                *config.Config
	currencyMenuItems  map[string]*systray.MenuItem
	updateTicker       *time.Ticker
	selectedCurrencies []string
)

func Initialize(cfg *config.Config) {
	selectedCurrencies = []string{"USD"}
	systray.SetIcon(utils.GetIcon(cfg.IconPath))
	systray.SetTitle("Currency App")
	systray.SetTooltip("Currency Exchange Rates")

	mRefresh := systray.AddMenuItem("Refresh Rates", "Get current exchange rates")
	mCurrencies := mRefresh.AddSubMenuItem("Select Currencies", "Choose currencies to track")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	currencyMenuItems = make(map[string]*systray.MenuItem)
	for _, curr := range cfg.Currencies {
		menuItem := mCurrencies.AddSubMenuItem(curr, "Toggle "+curr)
		currencyMenuItems[curr] = menuItem
		if curr == "USD" {
			menuItem.Check()
		}
	}

	go func() {
		for {
			select {
			case <-mRefresh.ClickedCh:
				refreshRates()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			default:
				handleCurrencySelection()
			}
		}
	}()

	startRateUpdater(cfg.UpdateInterval)
}

func handleCurrencySelection() {
	for curr, menuItem := range currencyMenuItems {
		select {
		case <-menuItem.ClickedCh:
			if menuItem.Checked() {
				menuItem.Uncheck()
				selectedCurrencies = removeString(selectedCurrencies, curr)
			} else {
				menuItem.Check()
				selectedCurrencies = append(selectedCurrencies, curr)
			}
			utils.LogInfo("Currency %s selection changed", curr)
			if len(selectedCurrencies) == 1 && updateTicker == nil {
				startRateUpdater(cfg.UpdateInterval)
			} else if len(selectedCurrencies) == 0 {
				stopRateUpdater()
			}
		default:
		}
	}
}

func startRateUpdater(dur int64) {
	updateTicker = time.NewTicker(time.Duration(dur) * time.Second)
	go func() {
		for range updateTicker.C {
			refreshRates()
		}
	}()
}

func stopRateUpdater() {
	if updateTicker != nil {
		updateTicker.Stop()
		updateTicker = nil
	}
}

func refreshRates() {
	rates, err := api.GetCurrentRates(cfg, selectedCurrencies)
	if err != nil {
		utils.LogError("Failed to get current rates: %v", err)
		return
	}

	err = storage.SaveRates(rates)
	if err != nil {
		utils.LogError("Failed to save rates: %v", err)
	}

	currency.UpdateTrayMenu(rates)
}

func removeString(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
