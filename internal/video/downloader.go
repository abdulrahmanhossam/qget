package video

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Download downloads a video from the given URL using yt-dlp.
// ytDlpPath is the full path to the yt-dlp executable.
// savePath is the directory where the video will be saved.
// denoPath is optional - if provided, deno should be in the same directory or PATH.
func Download(url string, ytDlpPath string, savePath string, denoPath string) error {
	output := filepath.Join(savePath, "%(title)s.%(ext)s")
	cmd := exec.Command(ytDlpPath, "-o", output, "--javascript-delay", "2000", url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}