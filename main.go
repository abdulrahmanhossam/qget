package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/abdulrahmanhossam/qget/internal/deps"
	"github.com/abdulrahmanhossam/qget/internal/ui"
	"github.com/abdulrahmanhossam/qget/internal/utils"
	"github.com/abdulrahmanhossam/qget/internal/video"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: qget <video-url>")
		os.Exit(1)
	}

	url := os.Args[1]

	found, ytDlpPath := deps.CheckYTDLP()
	if !found {
		fmt.Println("yt-dlp not found. Downloading...")
		if err := deps.DownloadYTDLP(); err != nil {
			fmt.Printf("Failed to download yt-dlp: %v\n", err)
			os.Exit(1)
		}
		if runtime.GOOS == "windows" {
			ytDlpPath = ".\\yt-dlp.exe"
		} else {
			ytDlpPath = "./yt-dlp"
		}
	}
	fmt.Printf("Using yt-dlp at: %s\n", ytDlpPath)

	denoFound, denoPath := deps.CheckDeno()
	if !denoFound {
		fmt.Println("deno not found. Downloading...")
		if err := deps.DownloadDeno(); err != nil {
			fmt.Printf("Failed to download deno: %v\n", err)
			os.Exit(1)
		}
		if runtime.GOOS == "windows" {
			denoPath = ".\\deno.exe"
		} else {
			denoPath = "./deno"
		}
	}
	fmt.Printf("Using deno at: %s\n", denoPath)

	savePath := utils.GetDownloadsDir()
	fmt.Printf("Saving to: %s\n", savePath)

	// Detect if URL contains a playlist (YouTube "list=" parameter).
	isPlaylist := strings.Contains(url, "list=")

	// Branch logic: playlist download vs single video download.
	if isPlaylist {
		// Ask user to confirm playlist download.
		confirm, err := ui.ConfirmPlaylist()
		if err != nil {
			fmt.Printf("Failed to confirm playlist: %v\n", err)
			os.Exit(1)
		}

		if confirm {
			// Download entire playlist at best quality (skip quality selection).
			fmt.Println("🚀 Starting playlist download (Best Quality)...")
			if err := video.Download(url, ytDlpPath, denoPath, savePath, "best", true); err != nil {
				fmt.Printf("Failed to download playlist: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Download complete!")
			return
		}
	}

	// Single video flow: fetch video info and let user select quality.
	fmt.Println("⏳ Fetching video qualities... Please wait.")
	info, err := video.GetVideoInfo(url, ytDlpPath, denoPath)
	if err != nil {
		fmt.Printf("Failed to get video info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Video Title:", info.Title)

	formatID, err := ui.SelectFormat(info.Formats)
	if err != nil {
		fmt.Printf("Failed to select format: %v\n", err)
		os.Exit(1)
	}

	if err := video.Download(url, ytDlpPath, denoPath, savePath, formatID, false); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download complete!")
}