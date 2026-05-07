package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/abdulrahmanhossam/qget/internal/deps"
	"github.com/abdulrahmanhossam/qget/internal/utils"
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

	// Check if yt-dlp is installed.
	found, ytDlpPath := deps.CheckYTDLP()
	if !found {
		fmt.Println("yt-dlp not found. Downloading...")
		if err := deps.DownloadYTDLP(); err != nil {
			fmt.Printf("❌ Failed to download yt-dlp: %v\n", err)
			os.Exit(1)
		}
		if runtime.GOOS == "windows" {
			ytDlpPath = ".\\yt-dlp.exe"
		} else {
			ytDlpPath = "./yt-dlp"
		}
	}
	fmt.Printf("Using yt-dlp at: %s\n", ytDlpPath)

	// Check if deno is installed.
	denoFound, _ := deps.CheckDeno()
	if !denoFound {
		fmt.Println("deno not found. Downloading...")
		if err := deps.DownloadDeno(); err != nil {
			fmt.Printf("❌ Failed to download deno: %v\n", err)
			os.Exit(1)
		}
		if runtime.GOOS == "windows" {
			fmt.Printf("Using deno at: .\\deno.exe\n")
		} else {
			fmt.Printf("Using deno at: ./deno\n")
		}
	}

	// Get the Downloads directory.
	savePath := utils.GetDownloadsDir()
	fmt.Printf("Saving to: %s\n", savePath)

	// Download the video.
	if err := video.Download(url, ytDlpPath, savePath); err != nil {
		fmt.Printf("❌ Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Download complete!")
}