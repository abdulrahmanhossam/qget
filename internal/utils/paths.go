package utils

import (
	"os"
	"path/filepath"
)

// GetAppDir returns the path to the application's data directory (~/.qget).
func GetAppDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(home, ".qget")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return appDir, nil
}
