//go:build windows
// +build windows

package autostart

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

const (
	runPath = `Software\Microsoft\Windows\CurrentVersion\Run`
	appName = "CurrencyApp"
)

func Enable() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, runPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	return key.SetStringValue(appName, filepath.ToSlash(exePath))
}

func Disable() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, runPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.DeleteValue(appName)
}
