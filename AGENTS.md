# qget

Go CLI that wraps yt-dlp + Deno + FFmpeg to download videos.

## Build

- `go build` — default binary (includes debug info)
- `go build -ldflags="-s -w" -o qget main.go` — stripped, small binary (used for releases)
- `go run .` — run without building
- Entrypoint is `main.go` at root (not under `cmd/`)

## Project layout

```
main.go                         # CLI entry, dependency init, routing
internal/deps/{checker,installer}.go
internal/ui/prompt.go            # survey/v2 prompts
internal/video/{info,downloader}.go
internal/utils/paths.go          # ~/.qget data dir + Downloads folder
```

## Key flow in main.go

1. Check yt-dlp, Deno, FFmpeg in `~/.qget` first, then PATH; auto-download if missing
2. URL from CLI arg or interactive survey prompt
3. Playlist detected by `list=` in URL. Confirm dialog, then either download full playlist or fall through to single-video path
4. Single video: fetch title, fetch resolutions via `--dump-json`, prompt quality+format, download
5. Playlist audio: download directly (no quality prompt). Playlist video: fetch first video's resolutions, prompt quality+format, download all

## Download args

- Always: `--newline --continue --js-runtimes deno:<path> --ffmpeg-location <path> --paths <dir> -o "%(title)s.%(ext)s" --no-playlist|--yes-playlist <url>`
- Audio: `-x --audio-format mp3 --audio-quality 0`
- Video: `-f <format> --merge-output-format <container>` where format is `bestvideo[height<=N]+bestaudio/best` for playlists with a specific quality, or `bestvideo+bestaudio/best` otherwise

## Prompts (survey/v2)

- Format: Video / Audio Only (MP3)
- Quality: sorted descending, prepended "Highest Available (Best)"
- Container: MP4 / MKV / WEBM

## Tests

- No test files exist in any package. No test command.

## Lint / format

- No lint config or CI. `gofmt` is the convention per CONTRIBUTING.md.
- No Makefile, no task runner config.

## Other notes

- `-o <dir>` flag for custom output directory (default `.`)
- Graceful interrupt handling via `os.Signal` goroutine
- The `qget` binary and `yt-dlp` are gitignored
- Uses Go 1.26.2, single module `github.com/abdulrahmanhossam/qget`
- GPL-3.0 license
