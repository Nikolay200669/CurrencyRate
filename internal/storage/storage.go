package storage

import (
	"fmt"
	"path/filepath"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
)

func SaveRates(rates []api.ExchangeRate) error {
	//for _, format := range cfg.SaveFormats {
	filename := filepath.Join("data", fmt.Sprintf("rates.%s", "csv"))

	//switch format {
	//case "json":
	//	if err := SaveMonthlyRatesToJSON(rates, filename); err != nil {
	//		return fmt.Errorf("failed to save monthly rates to JSON: %v", err)
	//	}
	//case "csv":
	if err := SaveRatesToCSV(rates, filename); err != nil {
		return fmt.Errorf("failed to save monthly rates to CSV: %v", err)
	}
	//case "xml":
	//	if err := SaveMonthlyRatesToXML(rates, filename); err != nil {
	//		return fmt.Errorf("failed to save monthly rates to XML: %v", err)
	//	}
	//default:
	//	return fmt.Errorf("unsupported format: %s", format)
	//}
	//}
	return nil
}
