package utils

import (
	"os"
	"path/filepath"
)

// GetDownloadsDir returns the path to the user's Downloads folder.
func GetDownloadsDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Downloads")
}