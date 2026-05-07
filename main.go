package main

import (
	"fmt"

	"github.com/abdulrahmanhossam/qget/internal/deps"
)

func main() {
	found, path := deps.CheckYTDLP()
	if found {
		fmt.Printf("✅ yt-dlp found at: %s\n", path)
	} else {
		fmt.Println("❌ yt-dlp is missing. Please install it first.")
	}
}