package video

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Download downloads a video from the given URL using yt-dlp.
// ytDlpPath is the full path to the yt-dlp executable.
// savePath is the directory where the video will be saved.
// denoPath is the path to the deno executable for JS processing.
func Download(url string, ytDlpPath string, savePath string, denoPath string) error {
	output := filepath.Join(savePath, "%(title)s.%(ext)s")
	cmd := exec.Command(ytDlpPath, "--js-runtimes", "deno:"+denoPath, "-o", output, url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}