package ui

import (
	"fmt"
	"sort"

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
		Options: []string{"Video", "Audio Only (MP3)"},
	}, &selected)
	if err != nil {
		return "", err
	}
	if selected == "Video" {
		return "video", nil
	}
	return "audio", nil
}

// SelectQuality prompts the user to select a download quality from available resolutions.
func SelectQuality(resolutions []int) (string, error) {
	if len(resolutions) == 0 {
		return "best", nil
	}

	sorted := make([]int, len(resolutions))
	copy(sorted, resolutions)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	options := []string{"Highest Available (Best)"}
	for _, h := range sorted {
		options = append(options, fmt.Sprintf("%dp", h))
	}

	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select download quality:",
		Options: options,
	}, &selected)
	if err != nil {
		return "", err
	}

	if selected == "Highest Available (Best)" {
		return "best", nil
	}
	for _, h := range sorted {
		if selected == fmt.Sprintf("%dp", h) {
			return fmt.Sprintf("%d", h), nil
		}
	}
	return "best", nil
}

// SelectVideoFormat prompts the user to select the final video container format.
func SelectVideoFormat() (string, error) {
	var selected string
	err := survey.AskOne(&survey.Select{
		Message: "Select final video format:",
		Options: []string{"MP4 (Most Compatible)", "MKV (Best for High Quality)", "WEBM (Open Source)"},
	}, &selected)
	if err != nil {
		return "", err
	}
	switch selected {
	case "MP4 (Most Compatible)":
		return "mp4", nil
	case "MKV (Best for High Quality)":
		return "mkv", nil
	case "WEBM (Open Source)":
		return "webm", nil
	}
	return "mp4", nil
}
