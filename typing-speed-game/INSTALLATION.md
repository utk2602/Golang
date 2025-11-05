# Installation Guide for Typing Speed Game CLI

## For Users Who Want to Install Your CLI

### Prerequisites
- Go 1.18 or higher installed
- Git (optional, for cloning)

### Method 1: Direct Install from GitHub (Easiest) ✅

Once you push your code to GitHub, anyone can install it with:

```bash
go install github.com/utk2602/Golang/typing-speed-game/cmd@latest
```

This will:
- Download the source code
- Build the binary
- Install it to `$GOPATH/bin` (usually `~/go/bin` on Linux/Mac or `%USERPROFILE%\go\bin` on Windows)

### Method 2: Install from Source

```bash
# Clone the repository
git clone https://github.com/utk2602/typing-speed-game.git
cd typing-speed-game

# Install
go install ./cmd
```

### Method 3: Build and Copy Manually

```bash
# Clone the repository
git clone https://github.com/utk2602/typing-speed-game.git
cd typing-speed-game

# Build the binary
go build -o typing-speed-game ./cmd/main.go

# Move to a directory in your PATH
# On Linux/Mac:
sudo mv typing-speed-game /usr/local/bin/

# On Windows (PowerShell as Admin):
Move-Item typing-speed-game.exe C:\Windows\System32\
```

## Setting Up PATH

### Windows

**Option 1: Temporary (current session only)**
```powershell
$env:Path += ";$env:USERPROFILE\go\bin"
```

**Option 2: Permanent**
1. Open System Properties > Environment Variables
2. Edit the `Path` variable for your user
3. Add `%USERPROFILE%\go\bin`
4. Click OK and restart your terminal

**Option 3: PowerShell Profile (Permanent)**
```powershell
# Edit your PowerShell profile
notepad $PROFILE

# Add this line:
$env:Path += ";$env:USERPROFILE\go\bin"

# Save and reload
. $PROFILE
```

### Linux/Mac

**Option 1: Temporary (current session only)**
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

**Option 2: Permanent**

Add to `~/.bashrc`, `~/.zshrc`, or `~/.profile`:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Then reload:
```bash
source ~/.bashrc  # or ~/.zshrc or ~/.profile
```

## Verify Installation

```bash
# Check if Go bin directory exists
go env GOPATH

# Try running the command
cmd

# Or check if it's in PATH
which cmd      # Linux/Mac
where.exe cmd  # Windows
```

## Troubleshooting

### "command not found" or "not recognized"

1. **Check if binary exists:**
   ```bash
   # Linux/Mac
   ls $(go env GOPATH)/bin/cmd
   
   # Windows
   dir %USERPROFILE%\go\bin\cmd.exe
   ```

2. **Check PATH:**
   ```bash
   # Linux/Mac
   echo $PATH | grep go
   
   # Windows
   echo $env:Path
   ```

3. **Reinstall:**
   ```bash
   go clean -i github.com/utk2602/Golang/typing-speed-game/cmd
   go install github.com/utk2602/Golang/typing-speed-game/cmd@latest
   ```

### Permission Denied (Linux/Mac)

```bash
chmod +x $(go env GOPATH)/bin/cmd
```

### Module Not Found

Make sure you have the latest Go version:
```bash
go version  # Should be 1.18 or higher
```

## Uninstalling

```bash
# Remove the binary
rm $(go env GOPATH)/bin/cmd        # Linux/Mac
del %USERPROFILE%\go\bin\cmd.exe  # Windows

# Remove module cache (optional)
go clean -modcache
```

## For Developers

If you're developing or contributing:

```bash
# Clone and build
git clone https://github.com/utk2602/typing-speed-game.git
cd typing-speed-game
go mod tidy
go build -o typing-speed-game ./cmd/main.go

# Run directly without installing
go run cmd/main.go

# Run tests
go test ./...
```
