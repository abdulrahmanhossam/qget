package video

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Download downloads a video from the given URL using yt-dlp.
// It accepts formatID to select specific video quality.
func Download(url string, ytDlpPath string, denoPath string, savePath string, formatID string) error {
	cmd := exec.Command(
		ytDlpPath,
		"--js-runtimes", "deno:"+denoPath,
		"-f", formatID+"+bestaudio/best",
		"-o", filepath.Join(savePath, "%(title)s.%(ext)s"),
		url,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}