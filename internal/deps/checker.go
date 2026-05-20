package deps

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdulrahmanhossam/qget/internal/utils"
)

// CheckYTDLP checks if yt-dlp is available in ~/.qget or system PATH.
func CheckYTDLP() (bool, string) {
	binaryName := "yt-dlp"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	appDir, err := utils.GetAppDir()
	if err == nil {
		localPath := filepath.Join(appDir, binaryName)
		if _, err := os.Stat(localPath); err == nil {
			absPath, err := filepath.Abs(localPath)
			if err == nil {
				return true, absPath
			}
			return true, localPath
		}
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}

// CheckDeno checks if deno is available in ~/.qget or system PATH.
func CheckDeno() (bool, string) {
	binaryName := "deno"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	appDir, err := utils.GetAppDir()
	if err == nil {
		localPath := filepath.Join(appDir, binaryName)
		if _, err := os.Stat(localPath); err == nil {
			absPath, err := filepath.Abs(localPath)
			if err == nil {
				return true, absPath
			}
			return true, localPath
		}
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}

// CheckFFmpeg checks if ffmpeg is available in ~/.qget or system PATH.
func CheckFFmpeg() (bool, string) {
	binaryName := "ffmpeg"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	appDir, err := utils.GetAppDir()
	if err == nil {
		localPath := filepath.Join(appDir, binaryName)
		if _, err := os.Stat(localPath); err == nil {
			absPath, err := filepath.Abs(localPath)
			if err == nil {
				return true, absPath
			}
			return true, localPath
		}
	}

	path, err := exec.LookPath(binaryName)
	if err == nil {
		return true, path
	}

	return false, ""
}
