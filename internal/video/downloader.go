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

var downloadProgressRegex = regexp.MustCompile(`\[download\]\s+([0-9]{1,3}\.[0-9])%`)

var playlistIndexRegex = regexp.MustCompile(`\[download\]\s+Downloading\s+video\s+([0-9]+)\s+of\s+([0-9]+)`)

func printProgress(percent float64, width int) {
	filled := int(float64(width) * percent / 100)
	var bar strings.Builder
	bar.Grow(width * 3)
	bar.WriteString(strings.Repeat("█", filled))
	bar.WriteString(strings.Repeat("-", width-filled))
	fmt.Printf("\r⏳ Downloading: [%s] %5.1f%%", bar.String(), percent)
}

func Download(url string, ytDlpPath string, denoPath string, ffmpegPath string, savePath string, formatID string, isPlaylist bool, isAudio bool) error {
	playlistFlag := "--no-playlist"
	if isPlaylist {
		playlistFlag = "--yes-playlist"
	}

	args := []string{
		"--newline",
		"--js-runtimes", "deno:" + denoPath,
		"--ffmpeg-location", ffmpegPath,
		"-o", filepath.Join(savePath, "%(title)s.%(ext)s"),
		playlistFlag,
		url,
	}

	if isAudio {
		args = append(args, "-x", "--audio-format", "mp3", "--audio-quality", "0")
	} else {
		formatArg := buildFormatArg(formatID, isPlaylist)
		args = append(args, "-f", formatArg)
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
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
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

func progressStep(symbol string, current int) string {
	if current == 0 {
		return ""
	}
	return symbol
}

// buildFormatArg constructs the yt-dlp format argument based on quality and playlist mode.
func buildFormatArg(formatID string, isPlaylist bool) string {
	if isPlaylist && formatID != "best" {
		return fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best", formatID)
	}
	return "bestvideo+bestaudio/best"
}