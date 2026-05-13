package deps

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// CheckYTDLP checks if yt-dlp is available in the current directory or system PATH.
func CheckYTDLP() (bool, string) {
	binaryName := "yt-dlp"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	if _, err := os.Stat(binaryName); err == nil {
		if runtime.GOOS == "windows" {
			return true, binaryName
		}
		return true, "./" + binaryName
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}

// CheckDeno checks if deno is available in the current directory or system PATH.
func CheckDeno() (bool, string) {
	binaryName := "deno"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	if _, err := os.Stat(binaryName); err == nil {
		if runtime.GOOS == "windows" {
			return true, binaryName
		}
		return true, "./" + binaryName
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}

// CheckFFmpeg checks if ffmpeg is available in the current directory or system PATH.
func CheckFFmpeg() (bool, string) {
	binaryName := "ffmpeg"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	if _, err := os.Stat(binaryName); err == nil {
		absPath, err := filepath.Abs(binaryName)
		if err == nil {
			return true, absPath
		}
		if runtime.GOOS == "windows" {
			return true, binaryName
		}
		return true, "./" + binaryName
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}

// checkBinary checks if a binary exists in the current directory first, then in PATH.
func checkBinary(binaryName string) (bool, string) {
	// Check current working directory first.
	cwd := filepath.Join(".", binaryName)
	if _, err := os.Stat(cwd); err == nil {
		return true, cwd
	}

	// Fallback to system PATH.
	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}