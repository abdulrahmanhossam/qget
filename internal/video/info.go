package video

import (
	"os/exec"
	"strings"
)

// GetVideoTitle fetches the video title using yt-dlp --print for fast extraction.
func GetVideoTitle(url string, ytDlpPath string, denoPath string) (string, error) {
	cmd := exec.Command(ytDlpPath, "--js-runtimes", "deno:"+denoPath, "--print", "title", "--no-playlist", url)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}