package deps

import (
	"os/exec"
	"runtime"
)

// CheckYTDLP checks if yt-dlp is available in the system PATH.
// Returns true and the full path if found, or false and empty string if not found.
func CheckYTDLP() (bool, string) {
	// Determine the correct binary name based on the operating system.
	// Windows uses .exe extension, Linux/macOS does not.
	var binaryName string
	if runtime.GOOS == "windows" {
		binaryName = "yt-dlp.exe"
	} else {
		binaryName = "yt-dlp"
	}

	// Use exec.LookPath to find the binary in the system PATH.
	// It returns the full path if found, or an error if not found.
	path, err := exec.LookPath(binaryName)
	if err != nil {
		return false, ""
	}

	return true, path
}