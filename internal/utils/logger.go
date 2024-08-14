package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLoggers(logDir string) error {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("currency_app_%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

func LogInfo(format string, v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
	}
}

func LogError(format string, v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
	}
}

func MonitorRateChanges(oldRate, newRate float64, currency string, threshold float64) {
	change := (newRate - oldRate) / oldRate
	if abs(change) > threshold {
		LogInfo("Significant change in %s rate: %.2f%% (old: %.4f, new: %.4f)",
			currency, change*100, oldRate, newRate)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
