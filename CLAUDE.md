# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

qget is a Go-based CLI video downloader that simplifies downloading videos from YouTube and other platforms. It auto-detects and installs required dependencies (yt-dlp, Deno, FFmpeg) on first run.

## Build Commands

```bash
go build          # Build the binary
go run .          # Run without building
```

## Architecture

The codebase is organized in a flat structure with packages under `internal/`:

- **main.go**: Entry point handling CLI arguments, dependency initialization, and download orchestration. Routes between single video and playlist modes.
- **internal/deps/**: Dependency management
  - `checker.go`: Checks if yt-dlp, Deno, FFmpeg exist in `~/.qget` or system PATH
  - `installer.go`: Downloads dependencies from GitHub releases to `~/.qget`
- **internal/ui/**: Interactive prompts using `survey/v2` library
  - `prompt.go`: Format selection (video/audio), quality selection, container format (MP4/MKV/WEBM)
- **internal/video/**: Video operations
  - `info.go`: Fetches video metadata and available resolutions via yt-dlp JSON output
  - `downloader.go`: Executes yt-dlp with progress parsing, handles format merging with FFmpeg
- **internal/utils/**: Path utilities
  - `paths.go`: Returns user's Downloads folder and creates `~/.qget` app directory

## Key Implementation Details

- **yt-dlp integration**: Uses `--js-runtimes deno:<path>` for JavaScript-based video extraction
- **FFmpeg merging**: Uses `--merge-output-format` for container conversion (MP4/MKV/WEBM)
- **Audio extraction**: Uses `-x --audio-format mp3` for MP3 downloads
- **Progress parsing**: Regex-based parsing of yt-dlp stdout for progress bar display
- **Platform-aware**: Handles Windows vs Linux binary names and extraction (zip vs tar.xz)

## Workflow

1. Check for dependencies in `~/.qget` first, then system PATH
2. Prompt for URL (CLI arg or interactive)
3. Detect playlist vs single video from URL (`list=` param)
4. Ask format type (video/audio)
5. Fetch available resolutions via `yt-dlp --dump-json`
6. Prompt quality and container format selection
7. Execute download with progress tracking