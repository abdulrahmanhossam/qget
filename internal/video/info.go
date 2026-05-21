package video

import (
	"encoding/json"
	"os/exec"
	"sort"
	"strings"
)

// Format represents a video format from yt-dlp JSON output.
type Format struct {
	Height int    `json:"height"`
	Ext    string `json:"ext"`
}

// VideoInfo represents video metadata from yt-dlp.
type VideoInfo struct {
	Title   string   `json:"title"`
	Formats []Format `json:"formats"`
}

// GetVideoTitle fetches the video title using yt-dlp --print for fast extraction.
func GetVideoTitle(url string, ytDlpPath string, denoPath string) (string, error) {
	cmd := exec.Command(ytDlpPath, "--js-runtimes", "deno:"+denoPath, "--print", "title", "--no-playlist", url)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// GetAvailableResolutions fetches available video heights from yt-dlp JSON.
// For playlists, it only fetches the first video to avoid freezing.
func GetAvailableResolutions(url string, ytDlpPath string, denoPath string, isPlaylist bool) ([]int, error) {
	args := []string{
		"--js-runtimes", "deno:" + denoPath,
		"--dump-json",
		"--no-download",
	}

	if isPlaylist {
		args = append(args, "--playlist-items", "1")
	} else {
		args = append(args, "--no-playlist")
	}

	args = append(args, url)

	cmd := exec.Command(ytDlpPath, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var info VideoInfo
	if err := json.Unmarshal(output, &info); err != nil {
		return nil, err
	}

	seen := map[int]bool{}
	var resolutions []int
	for _, f := range info.Formats {
		if f.Height > 0 && !seen[f.Height] {
			seen[f.Height] = true
			resolutions = append(resolutions, f.Height)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(resolutions)))
	return resolutions, nil
}
