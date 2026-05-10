package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abdulrahmanhossam/qget/internal/video"
)

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
