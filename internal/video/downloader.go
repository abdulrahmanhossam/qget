package video

import (
	"os"
	"os/exec"
)

// Download downloads a video from the given URL using yt-dlp.
// ytDlpPath is the full path to the yt-dlp executable.
func Download(url string, ytDlpPath string) error {
	// Create a new command that runs yt-dlp with the given URL.
	cmd := exec.Command(ytDlpPath, url)

	// Set stdout and stderr to the terminal so progress is visible.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and return any error that occurs.
	return cmd.Run()
}