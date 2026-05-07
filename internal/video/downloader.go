package video

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Download downloads a video from the given URL using yt-dlp.
// ytDlpPath is the full path to the yt-dlp executable.
// savePath is the directory where the video will be saved.
func Download(url string, ytDlpPath string, savePath string) error {
	// Create the yt-dlp command with output flag.
	output := filepath.Join(savePath, "%(title)s.%(ext)s")
	cmd := exec.Command(ytDlpPath, "-o", output, url)

	// Set stdout and stderr to the terminal so progress is visible.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and return any error that occurs.
	return cmd.Run()
}