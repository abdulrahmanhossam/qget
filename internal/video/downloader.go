package video

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Download downloads a video or playlist from the given URL using yt-dlp.
// If isPlaylist is true, it downloads the entire playlist; otherwise a single video.
// The formatID selects the video quality (e.g., "best", "1080p", "720p").
// The ffmpegPath specifies the path to the ffmpeg binary for post-processing.
func Download(url string, ytDlpPath string, denoPath string, ffmpegPath string, savePath string, formatID string, isPlaylist bool) error {
	playlistFlag := "--no-playlist"
	if isPlaylist {
		playlistFlag = "--yes-playlist"
	}

	args := []string{
		"--js-runtimes", "deno:" + denoPath,
		"--ffmpeg-location", ffmpegPath,
		"-f", formatID + "+bestaudio/best",
		"-o", filepath.Join(savePath, "%(title)s.%(ext)s"),
		playlistFlag,
		url,
	}

	cmd := exec.Command(ytDlpPath, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}