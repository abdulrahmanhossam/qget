# qget - CLI Video Downloader

A Go-based CLI tool that simplifies video downloading with zero manual setup.

## Key Features

- **Auto-detection and installation** of `yt-dlp`, `Deno`, and `FFmpeg`
- **Interactive Quality Selection** - choose your preferred video format
- **Intelligent Playlist Support** - download single videos or entire playlists
- **Cross-platform support** (Linux/Windows)
- **Interactive Mode** - run without arguments for a guided experience

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
2. Run these commands to make it globally executable:

```bash
chmod +x qget-linux
sudo mv qget-linux /usr/local/bin/qget
```

3. Now simply type `qget` from anywhere in your terminal.

## Usage

```bash
# Interactive mode - prompts for URL
qget

# Direct download
qget "https://youtube.com/watch?v=..."
```

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details.