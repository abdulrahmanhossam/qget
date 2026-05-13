package deps

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// DownloadYTDLP automatically downloads the yt-dlp binary from GitHub.
func DownloadYTDLP() error {
	var url string
	var filename string

	if runtime.GOOS == "windows" {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"
		filename = "yt-dlp.exe"
	} else {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"
		filename = "yt-dlp"
	}

	fmt.Printf("Downloading yt-dlp from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download yt-dlp: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download yt-dlp: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save yt-dlp: %w", err)
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(filename, 0755); err != nil {
			return fmt.Errorf("failed to set permissions: %w", err)
		}
	}

	fmt.Printf("yt-dlp downloaded successfully as: %s\n", filename)
	return nil
}

func DownloadDeno() error {
	var url string
	if runtime.GOOS == "windows" {
		url = "https://github.com/denoland/deno/releases/latest/download/deno-x86_64-pc-windows-msvc.zip"
	} else {
		url = "https://github.com/denoland/deno/releases/latest/download/deno-x86_64-unknown-linux-gnu.zip"
	}

	fmt.Printf("Downloading deno from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download deno: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download deno: HTTP %d", resp.StatusCode)
	}

	tmpFile, err := os.CreateTemp("", "deno-*.zip")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save deno zip: %w", err)
	}
	tmpFile.Close()

	zipFile, err := zip.OpenReader(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		if file.Name == "deno" || file.Name == "deno.exe" {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("failed to open zip entry: %w", err)
			}
			defer rc.Close()

			outFile, err := os.Create(file.Name)
			if err != nil {
				return fmt.Errorf("failed to create deno: %w", err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return fmt.Errorf("failed to extract deno: %w", err)
			}

			if runtime.GOOS != "windows" {
				os.Chmod(file.Name, 0755)
			}

			fmt.Printf("deno extracted successfully\n")
			return nil
		}
	}

	return fmt.Errorf("deno binary not found in zip")
}

// DownloadFFmpeg automatically downloads the ffmpeg binary from GitHub.
func DownloadFFmpeg() (string, error) {
	var url string
	var filename string

	if runtime.GOOS == "windows" {
		url = "https://github.com/yt-dlp/FFmpeg-Builds/releases/latest/download/ffmpeg-master-latest-win64-gpl.zip"
		filename = "ffmpeg.zip"
	} else {
		url = "https://github.com/yt-dlp/FFmpeg-Builds/releases/latest/download/ffmpeg-master-latest-linux64-gpl.tar.xz"
		filename = "ffmpeg.tar.xz"
	}

	fmt.Printf("Downloading FFmpeg from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download FFmpeg: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download FFmpeg: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save FFmpeg: %w", err)
	}
	file.Close()

	fmt.Println("Download complete. Extracting...")

	if runtime.GOOS == "windows" {
		if err := extractFFmpegWindows(filename); err != nil {
			return "", err
		}
	} else {
		if err := extractFFmpegLinux(filename); err != nil {
			return "", err
		}
	}

	os.Remove(filename)
	fmt.Println("FFmpeg extracted successfully")

	ffmpegPath := "./ffmpeg"
	if runtime.GOOS == "windows" {
		ffmpegPath = ".\\ffmpeg.exe"
	}
	absPath, err := filepath.Abs(ffmpegPath)
	if err != nil {
		return ffmpegPath, nil
	}
	return absPath, nil
}

func extractFFmpegWindows(zipPath string) error {
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		if filepath.Base(file.Name) == "ffmpeg.exe" && filepath.Dir(file.Name) == "bin" {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("failed to open zip entry: %w", err)
			}
			defer rc.Close()

			outFile, err := os.Create("ffmpeg.exe")
			if err != nil {
				return fmt.Errorf("failed to create ffmpeg.exe: %w", err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return fmt.Errorf("failed to extract ffmpeg.exe: %w", err)
			}

			fmt.Println("Extracted ffmpeg.exe to current directory")
			return nil
		}
	}

	return fmt.Errorf("ffmpeg.exe not found in zip")
}

func extractFFmpegLinux(tarPath string) error {
	cmd := exec.Command("tar", "-xf", tarPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tar: %w", err)
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var extractedDir string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "ffmpeg-") {
			extractedDir = entry.Name()
			break
		}
	}

	if extractedDir == "" {
		return fmt.Errorf("extracted ffmpeg directory not found")
	}

	ffmpegBin := filepath.Join(extractedDir, "bin", "ffmpeg")
	if _, err := os.Stat(ffmpegBin); err != nil {
		return fmt.Errorf("ffmpeg binary not found in extracted folder: %w", err)
	}

	if err := os.Rename(ffmpegBin, "./ffmpeg"); err != nil {
		return fmt.Errorf("failed to move ffmpeg: %w", err)
	}

	os.Chmod("./ffmpeg", 0755)

	os.RemoveAll(extractedDir)

	fmt.Println("Extracted ffmpeg to current directory")
	return nil
}

