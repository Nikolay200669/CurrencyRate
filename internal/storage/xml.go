package storage

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/Nikolay200669/CurrencyRate/internal/api"
)

type XMLRates struct {
	XMLName xml.Name          `xml:"rates"`
	Date    string            `xml:"date,attr"`
	Rates   []XMLExchangeRate `xml:"rate"`
}

type XMLExchangeRate struct {
	Currency string `xml:"currency,attr"`
	Buy      string `xml:"buy"`
	Sale     string `xml:"sale"`
}

func SaveRatesToXML(rates []api.ExchangeRate, filename string) error {
	xmlRates := XMLRates{
		Date:  time.Now().Format("2006-01-02"),
		Rates: make([]XMLExchangeRate, len(rates)),
	}

	for i, rate := range rates {
		xmlRates.Rates[i] = XMLExchangeRate{
			Currency: rate.Currency,
			Buy:      rate.Buy,
			Sale:     rate.Sale,
		}
	}

	xmlData, err := xml.MarshalIndent(xmlRates, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling XML: %v", err)
	}

	xmlData = append([]byte(xml.Header), xmlData...)

	err = os.WriteFile(filename, xmlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing XML file: %v", err)
	}

	return nil
}

func SaveMonthlyRatesToXML(rates map[string][]api.ExchangeRate, filename string) error {
	var monthlyRates struct {
		XMLName xml.Name   `xml:"monthly_rates"`
		Days    []XMLRates `xml:"day"`
	}

	for date, dailyRates := range rates {
		xmlRates := XMLRates{
			Date:  date,
			Rates: make([]XMLExchangeRate, len(dailyRates)),
		}

		for i, rate := range dailyRates {
			xmlRates.Rates[i] = XMLExchangeRate{
				Currency: rate.Currency,
				Buy:      rate.Buy,
				Sale:     rate.Sale,
			}
		}

		monthlyRates.Days = append(monthlyRates.Days, xmlRates)
	}

	xmlData, err := xml.MarshalIndent(monthlyRates, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling XML: %v", err)
	}

	xmlData = append([]byte(xml.Header), xmlData...)

	err = os.WriteFile(filename, xmlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing XML file: %v", err)
	}

	return nil
}
