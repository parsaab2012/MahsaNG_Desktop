//go:build windows

package ui

import (
	"fmt"
	"os/exec"
	"strings"

	"mahsang/internal/tun"
)

func wifiRepairSupported() bool {
	return true
}

func repairWiFi() error {
	if !tun.Elevated() {
		return fmt.Errorf("Wi-Fi repair requires Administrator privileges")
	}

	const script = `$ErrorActionPreference = 'Stop'
$adapters = @(Get-NetAdapter -Physical | Where-Object {
    $_.Name -match '(?i)wi-?fi|wlan' -or
    $_.InterfaceDescription -match '(?i)wi-?fi|wireless|802\.11|wlan'
})
if ($adapters.Count -ne 1) {
    throw "Expected exactly one physical Wi-Fi adapter; found $($adapters.Count)."
}
Restart-NetAdapter -Name $adapters[0].Name -Confirm:$false -ErrorAction Stop`

	cmd := exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		script,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		detail := strings.TrimSpace(string(output))
		if detail == "" {
			return fmt.Errorf("Restart-NetAdapter failed: %w", err)
		}
		return fmt.Errorf("Restart-NetAdapter failed: %w (%s)", err, detail)
	}

	return nil
}
