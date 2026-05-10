package video

import (
	"encoding/json"
	"os/exec"
)

type Format struct {
	FormatID  string `json:"format_id"`
	Resolution string `json:"resolution"`
	Ext       string `json:"ext"`
	Vcodec    string `json:"vcodec"`
	Height    int    `json:"height"`
}

type VideoInfo struct {
	Title   string   `json:"title"`
	Formats []Format `json:"formats"`
}

func GetVideoInfo(url string, ytDlpPath string, denoPath string) (*VideoInfo, error) {
	cmd := exec.Command(ytDlpPath, "--js-runtimes", "deno:"+denoPath, "--dump-json", "--no-playlist", url)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var videoInfo VideoInfo
	if err := json.Unmarshal(output, &videoInfo); err != nil {
		return nil, err
	}

	return &videoInfo, nil
}