package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/abdulrahmanhossam/qget/internal/deps"
	"github.com/abdulrahmanhossam/qget/internal/video"
)

func main() {
	// Check if the user provided a URL via command-line arguments.
	if len(os.Args) < 2 {
		fmt.Println("Usage: qget <video-url>")
		os.Exit(1)
	}

	// Get the URL from the first command-line argument.
	url := os.Args[1]

	// Check if yt-dlp is already installed.
	found, ytDlpPath := deps.CheckYTDLP()

	if !found {
		// If not found, download it automatically.
		fmt.Println("yt-dlp not found. Downloading...")

		if err := deps.DownloadYTDLP(); err != nil {
			fmt.Printf("❌ Failed to download yt-dlp: %v\n", err)
			os.Exit(1)
		}

		// Determine the path for the newly downloaded yt-dlp.
		if runtime.GOOS == "windows" {
			ytDlpPath = ".\\yt-dlp.exe"
		} else {
			ytDlpPath = "./yt-dlp"
		}
	}

	// Get the absolute path of the yt-dlp executable.
	absPath, err := filepath.Abs(ytDlpPath)
	if err != nil {
		absPath = ytDlpPath
	}

	fmt.Printf("Using yt-dlp at: %s\n", absPath)

	// Download the video.
	if err := video.Download(url, absPath); err != nil {
		fmt.Printf("❌ Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Download complete!")
}