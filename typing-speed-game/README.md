# Typing Speed Game

## Overview
The Typing Speed Game is a command-line interface (CLI) application designed to measure typing speed and provide analytics based on typing speed records and history. Players can engage in typing tests using predefined texts, and the application tracks their performance, offering insights into their typing speed over time.

## Features
- Measure typing speed in words per minute (WPM).
- Track individual game sessions and calculate typing speed.
- Analyze typing speed data to identify trends and averages.
- Store and retrieve typing speed records.
- User-friendly CLI interface for displaying prompts and results.

## Project Structure
```
typing-speed-game
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── game
│   │   ├── game.go     # Main game logic
│   │   └── session.go   # Individual game session management
│   ├── analytics
│   │   ├── analytics.go # Typing speed data analysis
│   │   └── stats.go     # Data structures for statistics
│   ├── storage
│   │   ├── storage.go   # Data storage management
│   │   └── file.go      # File I/O operations
│   └── ui
│       ├── display.go   # Display functions for CLI
│       └── input.go     # User input handling
├── pkg
│   ├── models
│   │   └── record.go    # Data model for typing speed records
│   └── utils
│       └── timer.go     # Utility functions for timing
├── data
│   └── texts.go         # Predefined texts for typing tests
├── go.mod               # Go module definition
├── go.sum               # Module dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd typing-speed-game
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage
- Start the game and follow the prompts to begin typing.
- After completing a typing test, your typing speed will be displayed.
- Access analytics to view your typing speed history and trends.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.