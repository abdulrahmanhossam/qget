package video

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// downloadProgressRegex matches yt-dlp download progress lines.
// Matches patterns like "[download]   0.0%" or "[download] 100.0%"
var downloadProgressRegex = regexp.MustCompile(`\[download\]\s+([0-9]{1,3}\.[0-9])%`)

// playlistIndexRegex matches playlist video index lines like "[download] Downloading video 2 of 5"
var playlistIndexRegex = regexp.MustCompile(`\[download\]\s+Downloading\s+video\s+([0-9]+)\s+of\s+([0-9]+)`)

// printProgress draws a single-line progress bar on the terminal.
func printProgress(percent float64, width int) {
	filled := int(float64(width) * percent / 100)
	empty := width - filled
	bar := strings.Repeat("█", filled) + strings.Repeat("-", empty)
	fmt.Printf("\r⏳ Downloading: [%s] %5.1f%%", bar, percent)
}

// Download downloads a video or playlist from the given URL using yt-dlp.
// If isPlaylist is true, it downloads the entire playlist; otherwise a single video.
// The formatID selects the video quality (e.g., "best", "1080p", "720p").
// The ffmpegPath specifies the path to the ffmpeg binary for post-processing.
// All raw yt-dlp output is hidden and replaced with a clean, single-line progress bar.
func Download(url string, ytDlpPath string, denoPath string, ffmpegPath string, savePath string, formatID string, isPlaylist bool) error {
	playlistFlag := "--no-playlist"
	if isPlaylist {
		playlistFlag = "--yes-playlist"
	}

	args := []string{
		"--newline",
		"--js-runtimes", "deno:" + denoPath,
		"--ffmpeg-location", ffmpegPath,
		"-f", formatID + "+bestaudio/best",
		"-o", filepath.Join(savePath, "%(title)s.%(ext)s"),
		playlistFlag,
		url,
	}

	cmd := exec.Command(ytDlpPath, args...)
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	barWidth := 30
	currentVideo := 0
	totalVideos := 0

	for scanner.Scan() {
		line := scanner.Text()

		if match := playlistIndexRegex.FindStringSubmatch(line); len(match) == 3 {
			currentVideo, _ = strconv.Atoi(match[1])
			totalVideos, _ = strconv.Atoi(match[2])
		}

		if match := downloadProgressRegex.FindStringSubmatch(line); len(match) == 2 {
			percent, _ := strconv.ParseFloat(match[1], 64)

			if isPlaylist && totalVideos > 0 {
				fmt.Printf("\r⏳ [%s %d/%d] ", progressStep("↓", currentVideo), currentVideo, totalVideos)
			}

			printProgress(percent, barWidth)
		}
	}

	fmt.Println()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// progressStep returns a filled step string based on current progress.
func progressStep(symbol string, current int) string {
	if current == 0 {
		return ""
	}
	return symbol
}