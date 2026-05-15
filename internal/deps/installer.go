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

	"github.com/abdulrahmanhossam/qget/internal/utils"
)

// DownloadYTDLP automatically downloads the yt-dlp binary from GitHub to ~/.qget.
func DownloadYTDLP() (string, error) {
	var url string
	var filename string

	if runtime.GOOS == "windows" {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"
		filename = "yt-dlp.exe"
	} else {
		url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"
		filename = "yt-dlp"
	}

	appDir, err := utils.GetAppDir()
	if err != nil {
		return "", fmt.Errorf("failed to get app directory: %w", err)
	}

	destPath := filepath.Join(appDir, filename)

	fmt.Printf("Downloading yt-dlp from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download yt-dlp: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download yt-dlp: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save yt-dlp: %w", err)
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(destPath, 0755); err != nil {
			return "", fmt.Errorf("failed to set permissions: %w", err)
		}
	}

	fmt.Printf("yt-dlp downloaded successfully to: %s\n", destPath)
	return destPath, nil
}

func DownloadDeno() (string, error) {
	var url string
	if runtime.GOOS == "windows" {
		url = "https://github.com/denoland/deno/releases/latest/download/deno-x86_64-pc-windows-msvc.zip"
	} else {
		url = "https://github.com/denoland/deno/releases/latest/download/deno-x86_64-unknown-linux-gnu.zip"
	}

	appDir, err := utils.GetAppDir()
	if err != nil {
		return "", fmt.Errorf("failed to get app directory: %w", err)
	}

	fmt.Printf("Downloading deno from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download deno: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download deno: HTTP %d", resp.StatusCode)
	}

	tmpFile, err := os.CreateTemp("", "deno-*.zip")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save deno zip: %w", err)
	}
	tmpFile.Close()

	zipFile, err := zip.OpenReader(tmpPath)
	if err != nil {
		return "", fmt.Errorf("failed to open zip: %w", err)
	}
	defer zipFile.Close()

	binaryName := "deno"
	if runtime.GOOS == "windows" {
		binaryName = "deno.exe"
	}
	destPath := filepath.Join(appDir, binaryName)

	for _, file := range zipFile.File {
		if file.Name == "deno" || file.Name == "deno.exe" {
			rc, err := file.Open()
			if err != nil {
				return "", fmt.Errorf("failed to open zip entry: %w", err)
			}
			defer rc.Close()

			outFile, err := os.Create(destPath)
			if err != nil {
				return "", fmt.Errorf("failed to create deno: %w", err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return "", fmt.Errorf("failed to extract deno: %w", err)
			}

			if runtime.GOOS != "windows" {
				os.Chmod(destPath, 0755)
			}

			fmt.Printf("deno extracted successfully to: %s\n", destPath)
			return destPath, nil
		}
	}

	return "", fmt.Errorf("deno binary not found in zip")
}

// DownloadFFmpeg automatically downloads the ffmpeg binary from GitHub to ~/.qget.
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

	appDir, err := utils.GetAppDir()
	if err != nil {
		return "", fmt.Errorf("failed to get app directory: %w", err)
	}

	destPath := filepath.Join(appDir, filename)

	fmt.Printf("Downloading FFmpeg from: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download FFmpeg: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download FFmpeg: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
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
		if err := extractFFmpegWindows(destPath, appDir); err != nil {
			return "", err
		}
	} else {
		if err := extractFFmpegLinux(destPath, appDir); err != nil {
			return "", err
		}
	}

	os.Remove(destPath)
	fmt.Println("FFmpeg extracted successfully")

	ffmpegPath := filepath.Join(appDir, "ffmpeg")
	if runtime.GOOS == "windows" {
		ffmpegPath = filepath.Join(appDir, "ffmpeg.exe")
	}
	return ffmpegPath, nil
}

func extractFFmpegWindows(zipPath, destDir string) error {
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

			outFile, err := os.Create(filepath.Join(destDir, "ffmpeg.exe"))
			if err != nil {
				return fmt.Errorf("failed to create ffmpeg.exe: %w", err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return fmt.Errorf("failed to extract ffmpeg.exe: %w", err)
			}

			fmt.Println("Extracted ffmpeg.exe to ~/.qget")
			return nil
		}
	}

	return fmt.Errorf("ffmpeg.exe not found in zip")
}

func extractFFmpegLinux(tarPath, destDir string) error {
	origDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer os.Chdir(origDir)

	tmpDir := filepath.Dir(tarPath)
	if err := os.Chdir(tmpDir); err != nil {
		return fmt.Errorf("failed to change to temp directory: %w", err)
	}

	cmd := exec.Command("tar", "-xf", tarPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tar: %w", err)
	}

	entries, err := os.ReadDir(tmpDir)
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

	ffmpegBin := filepath.Join(tmpDir, extractedDir, "bin", "ffmpeg")
	if _, err := os.Stat(ffmpegBin); err != nil {
		return fmt.Errorf("ffmpeg binary not found in extracted folder: %w", err)
	}

	if err := os.Rename(ffmpegBin, filepath.Join(destDir, "ffmpeg")); err != nil {
		return fmt.Errorf("failed to move ffmpeg: %w", err)
	}

	os.Chmod(filepath.Join(destDir, "ffmpeg"), 0755)

	os.RemoveAll(filepath.Join(tmpDir, extractedDir))

	fmt.Println("Extracted ffmpeg to ~/.qget")
	return nil
}