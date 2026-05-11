<h2 align="center">Zima-Command@v1.0</h2>

» Installation

**Linux** (in `/build` directory):
for install:
```bash
  bash ./install.sh 
```

**Windows** (Run as Administrator):

`install.bat`      # Install  
`uninstall.bat`    # Uninstall

## Build from Source (Optional)

`go build -o zc`

Note: Go compiler is only needed for building. The binary runs standalone without Go.

## Usage

After installation:

`zc --help`      # Show help page  
`zc info`        # Show system information  
`zc back`        # Go back one directory

## Available Commands (Beta)

- `zc info` - Display system information (OS, CPU, RAM, Disk, GPU, Timezone, Uptime)
- `zc back` - Navigate one directory back
- `zc --help` - Show help menu

More commands coming soon.

## System Information Displayed

When running `zc info`, you will see:
- Operating System and Architecture
- CPU Cores and Threads
- Hostname and User
- Current Date, Time, and Timezone
- System Uptime
- Total, Used, and Free RAM
- Total, Used, and Free Disk Space
- GPU Information (if available)

## Requirements

- Linux or Windows operating system
- No additional dependencies needed for running
- Go compiler required only for building from source

## Uninstall

**Linux:** `bash ./uninstall.sh`

**Windows:** Run `uninstall.bat` as Administrator

## Links

- Source: github.com/amodemoli/zc
- License: MIT
- Version: 0.1 (Beta)

## Notes

- Currently only Linux and Windows are supported
- For Windows, Windows Terminal is recommended for best color display
- Run as administrator on Windows for proper installation
