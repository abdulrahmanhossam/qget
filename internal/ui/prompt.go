package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abdulrahmanhossam/qget/internal/video"
)

// SelectFormat displays an interactive dropdown for users to choose a video format.
// It filters out non-video formats (audio-only, mhtml) and only shows mp4/webm.
func SelectFormat(formats []video.Format) (string, error) {
	options := []string{}
	optionToID := map[string]string{}

	for _, f := range formats {
		if f.Height <= 0 || (f.Ext != "mp4" && f.Ext != "webm") {
			continue
		}

		option := fmt.Sprintf("[%dp] %s (Codec: %s) - ID: %s", f.Height, f.Ext, f.Vcodec, f.FormatID)
		options = append(options, option)
		optionToID[option] = f.FormatID
	}

	if len(options) == 0 {
		return "", fmt.Errorf("no video formats available")
	}

	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select Video Quality:",
		Options: options,
	}, &selected)
	if err != nil {
		return "", err
	}

	return optionToID[selected], nil
}

// ConfirmPlaylist asks the user whether they want to download an entire playlist.
func ConfirmPlaylist() (bool, error) {
	var confirm bool
	err := survey.AskOne(&survey.Confirm{
		Message: "This URL contains a playlist. Do you want to download the entire playlist at the best quality?",
		Default: false,
	}, &confirm)
	if err != nil {
		return false, err
	}
	return confirm, nil
}

// AskFormatType prompts the user to choose between video or audio-only download.
func AskFormatType() (string, error) {
	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "What do you want to download?",
		Options: []string{"🎥 Video", "🎵 Audio Only (MP3)"},
	}, &selected)
	if err != nil {
		return "", err
	}
	if selected == "🎥 Video" {
		return "video", nil
	}
	return "audio", nil
}

// SelectPlaylistQuality prompts the user to select a uniform quality for playlist downloads.
func SelectPlaylistQuality() (string, error) {
	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select download quality for the entire playlist:",
		Options: []string{
			"🌟 Highest Available",
			"🎬 1080p",
			"📺 720p",
			"💿 480p",
			"📱 360p",
		},
	}, &selected)
	if err != nil {
		return "", err
	}
	switch selected {
	case "🌟 Highest Available":
		return "best", nil
	case "🎬 1080p":
		return "1080", nil
	case "📺 720p":
		return "720", nil
	case "💿 480p":
		return "480", nil
	case "📱 360p":
		return "360", nil
	}
	return "best", nil
}
