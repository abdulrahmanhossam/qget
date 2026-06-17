# qget - CLI Video Downloader

A minimalist, zero-setup CLI tool built in Go that wraps yt-dlp, Deno, and FFmpeg to download videos from YouTube and other platforms.

## Quick Install (Linux)

```bash
curl -sL https://raw.githubusercontent.com/abdulrahmanhossam/qget/main/install.sh | bash
```

## Features

- Auto-detection and installation of yt-dlp, Deno, and FFmpeg on first run
- Custom single-line progress bar for clean download visualization
- Audio-only extraction (MP3) for music and podcasts
- Format selection: MP4, MKV, or WEBM container
- Dynamic quality menu sorted from highest to lowest resolution
- Intelligent playlist support (single video or full playlist)
- Custom output directory via the `-o` flag
- Download resuming via yt-dlp's `--continue` flag
- Multi-dependency update via the `--update` flag
- Cross-platform support (Linux and Windows)

## Installation

### Windows

1. Download `qget-windows.exe` from the [Releases](https://github.com/abdulrahmanhossam/qget/releases) page.
2. Double-click to launch interactive mode, or use in Command Prompt:

```cmd
qget-windows.exe "https://youtube.com/watch?v=..."
```

### Linux

1. Download `qget-linux` from the [Releases](https://github.com/abdulrahmanhossam/qget/releases) page.
2. Make it executable and move it to your PATH:

```bash
chmod +x qget-linux
sudo mv qget-linux /usr/local/bin/qget
```

### Build from source

```bash
go build -ldflags="-s -w" -o qget main.go
```

## Usage

```bash
# Interactive mode - prompts for URL
qget

# Direct download
qget "https://youtube.com/watch?v=..."

# Custom output directory
qget -o ~/Videos "https://youtube.com/watch?v=..."

# Update dependencies (yt-dlp and Deno)
qget --update
```

## License

[GPL-3.0 License](LICENSE)
