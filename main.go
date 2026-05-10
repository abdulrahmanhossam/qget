package main

import (
	"fmt"
	"os"
	"runtime"

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

	if err := video.Download(url, ytDlpPath, denoPath, savePath, formatID); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download complete!")
}