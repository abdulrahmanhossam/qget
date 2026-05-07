package deps

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

// DownloadYTDLP automatically downloads the yt-dlp binary from GitHub.
// It determines the correct URL based on the operating system.
func DownloadYTDLP() error {
	// Determine the correct download URL and filename based on the OS.
	var url string
	var filename string

	if runtime.GOOS == "windows" {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"
		filename = "yt-dlp.exe"
	} else {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"
		filename = "yt-dlp"
	}

	fmt.Printf("Downloading yt-dlp from: %s\n", url)

	// Make a GET request to download the file.
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download yt-dlp: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download yt-dlp: HTTP %d", resp.StatusCode)
	}

	// Create the output file in the current working directory.
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Copy the downloaded content to the file.
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save yt-dlp: %w", err)
	}

	// Set executable permissions on Linux/macOS.
	if runtime.GOOS != "windows" {
		if err := os.Chmod(filename, 0755); err != nil {
			return fmt.Errorf("failed to set permissions: %w", err)
		}
	}

	fmt.Printf("✅ yt-dlp downloaded successfully as: %s\n", filename)
	return nil
}