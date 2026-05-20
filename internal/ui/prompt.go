package ui

import (
	"github.com/AlecAivazis/survey/v2"
)

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

// SelectQuality prompts the user to select a download quality for video content.
func SelectQuality() (string, error) {
	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select download quality:",
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

// SelectVideoFormat prompts the user to select the final video container format.
func SelectVideoFormat() (string, error) {
	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select final video format:",
		Options: []string{
			"🎞️ MP4 (Most Compatible)",
			"🎬 MKV (Best for High Quality)",
			"🌐 WEBM (Open Source)",
		},
	}, &selected)
	if err != nil {
		return "", err
	}
	switch selected {
	case "🎞️ MP4 (Most Compatible)":
		return "mp4", nil
	case "🎬 MKV (Best for High Quality)":
		return "mkv", nil
	case "🌐 WEBM (Open Source)":
		return "webm", nil
	}
	return "mp4", nil
}