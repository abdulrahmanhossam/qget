package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abdulrahmanhossam/qget/internal/deps"
	"github.com/abdulrahmanhossam/qget/internal/ui"
	"github.com/abdulrahmanhossam/qget/internal/utils"
	"github.com/abdulrahmanhossam/qget/internal/video"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n⚠️ Interrupted by user. Exiting gracefully...")
		os.Exit(0)
	}()

	var url string

	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Println("🔗 No URL provided. Launching interactive mode...")
		prompt := &survey.Input{
			Message: "🔗 Enter video or playlist URL:",
		}
		if err := survey.AskOne(prompt, &url); err != nil {
			fmt.Println("❌ No URL provided. Exiting...")
			os.Exit(0)
		}
		if url == "" {
			fmt.Println("❌ No URL provided. Exiting...")
			os.Exit(0)
		}
	}

	ytDlpFound, ytDlpPath := deps.CheckYTDLP()
	if !ytDlpFound {
		fmt.Println("yt-dlp not found. Downloading...")
		downloadedPath, err := deps.DownloadYTDLP()
		if err != nil {
			fmt.Printf("Failed to download yt-dlp: %v\n", err)
			os.Exit(1)
		}
		ytDlpPath = downloadedPath
	}
	fmt.Printf("Using yt-dlp at: %s\n", ytDlpPath)

	denoFound, denoPath := deps.CheckDeno()
	if !denoFound {
		fmt.Println("deno not found. Downloading...")
		downloadedPath, err := deps.DownloadDeno()
		if err != nil {
			fmt.Printf("Failed to download deno: %v\n", err)
			os.Exit(1)
		}
		denoPath = downloadedPath
	}
	fmt.Printf("Using deno at: %s\n", denoPath)

	ffmpegFound, ffmpegPath := deps.CheckFFmpeg()
	if !ffmpegFound {
		fmt.Println("FFmpeg not found. Downloading...")
		downloadedPath, err := deps.DownloadFFmpeg()
		if err != nil {
			fmt.Printf("Failed to download FFmpeg: %v\n", err)
			os.Exit(1)
		}
		ffmpegPath = downloadedPath
	}
	fmt.Printf("Using FFmpeg at: %s\n", ffmpegPath)

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
			formatType, err := ui.AskFormatType()
			if err != nil {
				fmt.Printf("Failed to get format type: %v\n", err)
				os.Exit(1)
			}

			if formatType == "audio" {
				fmt.Println("🎵 Starting Audio Download (MP3)...")
				if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, "", true, true, ""); err != nil {
					fmt.Printf("Failed to download playlist: %v\n", err)
					os.Exit(1)
				}
			} else {
				quality, err := ui.SelectQuality()
				if err != nil {
					fmt.Printf("Failed to select quality: %v\n", err)
					os.Exit(1)
				}

				containerFormat, err := ui.SelectVideoFormat()
				if err != nil {
					fmt.Printf("Failed to select format: %v\n", err)
					os.Exit(1)
				}

				fmt.Printf("🚀 Starting playlist download (Quality: %s, Format: %s)...\n", quality, containerFormat)
				if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, quality, true, false, containerFormat); err != nil {
					fmt.Printf("Failed to download playlist: %v\n", err)
					os.Exit(1)
				}
			}
			fmt.Println("Download complete!")
			return
		}

		// User declined playlist: check if this is a pure playlist URL (no video ID).
		hasVideoID := strings.Contains(url, "watch?v=") || strings.Contains(url, "youtu.be/")
		if !hasVideoID {
			fmt.Println("Error: This is a playlist-only URL. There is no single video to fetch qualities for.")
			os.Exit(1)
		}
	}

	// Ask user whether they want video or audio-only download.
	formatType, err := ui.AskFormatType()
	if err != nil {
		fmt.Printf("Failed to get format type: %v\n", err)
		os.Exit(1)
	}

	if formatType == "audio" {
		fmt.Println("🎵 Starting Audio Download (MP3)...")
		if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, "", false, true, ""); err != nil {
			fmt.Printf("Failed to download audio: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Download complete!")
		return
	}

	fmt.Println("⏳ Fetching video info...")
	title, err := video.GetVideoTitle(url, ytDlpPath, denoPath)
	if err != nil {
		fmt.Printf("Failed to get video title: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Video Title:", title)

	quality, err := ui.SelectQuality()
	if err != nil {
		fmt.Printf("Failed to select quality: %v\n", err)
		os.Exit(1)
	}

	containerFormat, err := ui.SelectVideoFormat()
	if err != nil {
		fmt.Printf("Failed to select format: %v\n", err)
		os.Exit(1)
	}

	if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, quality, false, false, containerFormat); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download complete!")
}