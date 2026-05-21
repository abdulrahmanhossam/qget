# qget - CLI Video Downloader

A Go-based CLI tool that simplifies video downloading with zero manual setup.

## Quick Install / Update (Linux)

```bash
curl -sL https://raw.githubusercontent.com/abdulrahmanhossam/qget/main/install.sh | bash
```

## Features

- **Auto-detection and installation** of `yt-dlp`, `Deno`, and `FFmpeg`
- **Interactive Quality Selection** - choose your preferred video format
- **Intelligent Playlist Support** - download single videos or entire playlists
- **Cross-platform support** (Linux/Windows)
- **Interactive Mode** - run without arguments for a guided experience
- **Minimalist single-line progress bar** for clean download visualization
- **Audio-only (MP3) extraction** for music and podcast downloads
- **Unified smart quality selection** for playlists and single videos
- **Custom output directories** via `-o` flag

## Installation

### 🪟 Windows Users

1. Download `qget-windows.exe` from the [Releases](https://github.com/abdulrahmanhossam/qget/releases) page.
2. **Double-click** to launch interactive mode, or use in CMD:

```cmd
qget-windows.exe "https://youtube.com/watch?v=..."
```

No installation or setup required!

### 🐧 Linux Users

1. Download `qget-linux` from the [Releases](https://github.com/abdulrahmanhossam/qget/releases) page.
2. Navigate to the folder where you downloaded the file (typically `~/Downloads`) and run:

```bash
# 1. Navigate to where you downloaded the file
cd ~/Downloads

# 2. Make it executable
chmod +x qget-linux

# 3. Move it to your global bin directory
sudo mv qget-linux /usr/local/bin/qget
```

3. Now simply type `qget` from anywhere in your terminal.

### Build from source

```bash
go build -ldflags="-s -w" -o qget main.go
```

The `-ldflags="-s -w"` strips debugging information for an extremely small binary.

## Usage

```bash
# Interactive mode - prompts for URL
qget

# Direct download
qget "https://youtube.com/watch?v=..."

# Custom output directory
qget -o ~/Videos "https://youtube.com/watch?v=..."
```

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details.