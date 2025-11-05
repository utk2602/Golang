# Typing Speed Game 🚀⌨️

A sleek command-line typing speed test application built in Go. Test your typing speed, track your progress, and improve your skills with detailed analytics!

## Features ✨

- ⚡ Real-time typing speed measurement (WPM)
- 📊 Detailed accuracy tracking
- 💾 Persistent record storage
- 📈 Comprehensive analytics (average, best, worst speeds)
- 📜 Game history with recent records
- 🎯 Multiple practice texts
- 🖥️ Beautiful CLI interface

## Installation 📦

### Option 1: Install directly from GitHub (Recommended)

```bash
go install github.com/utk2602/typing-speed-game/cmd@latest
```

This will install the `cmd` binary to your `$GOPATH/bin` directory.

### Option 2: Build from source

```bash
# Clone the repository
git clone https://github.com/utk2602/typing-speed-game.git
cd typing-speed-game

# Build and install
go install ./cmd

# Or build an executable
go build -o typing-speed-game ./cmd/main.go
```

### Make sure Go bin is in your PATH

**Windows (PowerShell):**
```powershell
$env:Path += ";$env:USERPROFILE\go\bin"
```

**Linux/Mac:**
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage 🎮

Simply run:
```bash
cmd
```

Or if you built with a custom name:
```bash
./typing-speed-game
```

### Menu Options:

1. **Start New Game** - Begin a new typing test
2. **View Analytics** - See your performance statistics
3. **Exit** - Quit the application

## How It Works 🔧

1. **Type a random text** - You'll be shown a random sentence to type
2. **Timing starts** - Timer begins when you start typing
3. **Calculate metrics** - WPM and accuracy are calculated automatically
4. **Save results** - Your record is saved to `typing_records.json`
5. **Track progress** - View your statistics and improvement over time

## Project Structure 📁

```
typing-speed-game
├── cmd
│   └── main.go          # CLI entry point
├── internal
│   ├── game
│   │   ├── game.go      # Core game logic
│   │   └── session.go   # Session management & calculations
│   ├── analytics
│   │   ├── analytics.go # Performance analytics
│   │   └── stats.go     # Statistics tracking
│   ├── storage
│   │   ├── storage.go   # Thread-safe record storage
│   │   └── file.go      # File I/O operations
│   └── ui
│       ├── display.go   # Display formatting
│       └── input.go     # User input handling
├── pkg
│   ├── models
│   │   └── record.go    # Data models
│   └── utils
│       └── timer.go     # Timing utilities
├── data
│   └── texts.go         # Practice text collection
└── go.mod               # Go module definition
```

## Development 🛠️

### Prerequisites
- Go 1.18 or higher

### Local Development

```bash
# Clone the repository
git clone https://github.com/utk2602/typing-speed-game.git
cd typing-speed-game

# Install dependencies
go mod tidy

# Run locally
go run cmd/main.go

# Run tests
go test ./...

# Build
go build -o typing-speed-game ./cmd/main.go
```

## Publishing Your CLI 📤

To make your CLI installable by others:

1. **Push to GitHub:**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/utk2602/typing-speed-game.git
   git push -u origin main
   ```

2. **Users can then install with:**
   ```bash
   go install github.com/utk2602/typing-speed-game/cmd@latest
   ```

## Data Storage 💾

Game records are stored in `typing_records.json` in your current directory:
```json
[
  {
    "player_name": "utkarsh",
    "speed": 37.70,
    "timestamp": "2025-11-05T10:30:00Z"
  }
]
```

## Contributing 🤝

Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests

## License 📄

This project is licensed under the MIT License.

## Author ✍️

**Utkarsh** - [@utk2602](https://github.com/utk2602)

## Acknowledgments 🙏

Built with ❤️ using Go

---

**Happy Typing! 🎉**