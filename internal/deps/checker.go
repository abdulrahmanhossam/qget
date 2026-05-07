package deps

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// CheckYTDLP checks if yt-dlp is available in the system PATH.
// Returns true and the full path if found, or false and empty string if not found.
func CheckYTDLP() (bool, string) {
	binaryName := "yt-dlp"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	return checkBinary(binaryName)
}

// CheckDeno checks if deno is available in the system PATH or current directory.
// Returns true and the full path if found, or false and empty string if not found.
func CheckDeno() (bool, string) {
	binaryName := "deno"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	return checkBinary(binaryName)
}

// checkBinary checks if a binary exists in PATH or current directory.
func checkBinary(binaryName string) (bool, string) {
	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	execPath, err := os.Executable()
	if err != nil {
		return false, ""
	}

	localPath := filepath.Join(filepath.Dir(execPath), binaryName)
	if _, err := os.Stat(localPath); err == nil {
		return true, localPath
	}

	return false, ""
}