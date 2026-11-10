//go:build !windows

package ui

import "errors"

func wifiRepairSupported() bool {
	return false
}

func repairWiFi() error {
	return errors.New("Wi-Fi recovery is only available on Windows")
}
