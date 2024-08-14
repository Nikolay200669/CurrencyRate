//go:build darwin
// +build darwin

package autostart

import (
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const (
	launchAgentsDir = "Library/LaunchAgents"
	plistFile       = "com.currencyapp.plist"
	plistTemplate   = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.currencyapp</string>
    <key>ProgramArguments</key>
    <array>
        <string>{{.ExecPath}}</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
`
)

func Enable() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	launchAgentsPath := filepath.Join(homeDir, launchAgentsDir)
	if err := os.MkdirAll(launchAgentsPath, 0755); err != nil {
		return err
	}

	plistPath := filepath.Join(launchAgentsPath, plistFile)
	file, err := os.Create(plistPath)
	if err != nil {
		return err
	}
	defer file.Close()

	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	tmpl, err := template.New("plist").Parse(plistTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, struct{ ExecPath string }{ExecPath: execPath}); err != nil {
		return err
	}

	return exec.Command("launchctl", "load", plistPath).Run()
}

func Disable() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	plistPath := filepath.Join(homeDir, launchAgentsDir, plistFile)

	if err := exec.Command("launchctl", "unload", plistPath).Run(); err != nil {
		return err
	}

	return os.Remove(plistPath)
}
