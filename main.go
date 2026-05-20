package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unicode"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abdulrahmanhossam/qget/internal/deps"
	"github.com/abdulrahmanhossam/qget/internal/ui"
	"github.com/abdulrahmanhossam/qget/internal/utils"
	"github.com/abdulrahmanhossam/qget/internal/video"
	"golang.org/x/text/unicode/bidi"
)

// sanitizeTitle removes non-printable, complex symbols, and emoji
// that break terminal rendering, keeping alphanumeric, spaces, and basic punctuation.
func sanitizeTitle(text string) string {
	var result strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '.' || r == '-' || r == '_' || r == '(' || r == ')' {
			result.WriteRune(r)
		}
	}
	return strings.TrimSpace(result.String())
}

// shapeArabicText handles RTL Arabic text for terminal display.
func shapeArabicText(text string) string {
	sanitized := sanitizeTitle(text)
	if sanitized == "" {
		return text
	}

	paragraph := &bidi.Paragraph{}
	if _, err := paragraph.SetString(sanitized); err != nil {
		return sanitized
	}

	ordering, err := paragraph.Order()
	if err != nil {
		return sanitized
	}

	var result strings.Builder
	for i := 0; i < ordering.NumRuns(); i++ {
		run := ordering.Run(i)
		result.WriteString(run.String())
	}
	return result.String()
}

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
			// Download entire playlist at best quality (skip quality selection).
			fmt.Println("🚀 Starting playlist download (Best Quality)...")
			if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, "best", true); err != nil {
				fmt.Printf("Failed to download playlist: %v\n", err)
				os.Exit(1)
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

	// Single video flow: fetch video info and let user select quality.
	fmt.Println("⏳ Fetching video qualities... Please wait.")
	info, err := video.GetVideoInfo(url, ytDlpPath, denoPath)
	if err != nil {
		fmt.Printf("Failed to get video info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Video Title:", shapeArabicText(info.Title))

	formatID, err := ui.SelectFormat(info.Formats)
	if err != nil {
		fmt.Printf("Failed to select format: %v\n", err)
		os.Exit(1)
	}

	if err := video.Download(url, ytDlpPath, denoPath, ffmpegPath, savePath, formatID, false); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download complete!")
}
