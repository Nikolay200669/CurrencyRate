//go:build linux
// +build linux

package autostart

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const (
	autostartDir = ".config/autostart"
	desktopEntry = "CurrencyApp.desktop"
	desktopTmpl  = `[Desktop Entry]
Type=Application
Name=CurrencyApp
Exec={{.ExecPath}}
Hidden=false
NoDisplay=false
X-GNOME-Autostart-enabled=true
`
)

func Enable() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	autostartPath := filepath.Join(homeDir, autostartDir)
	if err := os.MkdirAll(autostartPath, 0755); err != nil {
		return err
	}

	desktopFilePath := filepath.Join(autostartPath, desktopEntry)
	file, err := os.Create(desktopFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	tmpl, err := template.New("desktop").Parse(desktopTmpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, struct{ ExecPath string }{ExecPath: execPath})
}

func Disable() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	desktopFilePath := filepath.Join(homeDir, autostartDir, desktopEntry)
	if err := os.Remove(desktopFilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove autostart entry: %v", err)
	}

	return nil
}
