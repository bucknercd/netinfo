# netinfo

`netinfo` is a lightweight CLI utility written in Go that inspects and reports local network state, interfaces, and connectivity.

It provides a quick diagnostic snapshot of:
- Host identity
- Network interfaces
- Local IP configuration
- Active interface detection
- DNS resolution
- Internet connectivity
- Default gateway
- Public IP address

## Instructions for use

1. If you want the binary tool to use, browse the binaries directory and download the appropriate OS/Arch executable
2. Alternatively this module can be installed via go get: `go get github.com/bucknercd/netinfo` and then `go install` if you have Golang
3. Finally, you could clone this repository if you have Golang

## Features

- Lists all network interfaces with:
  - name
  - hardware address
  - status (UP/DOWN/RUNNING/LOOPBACK)
  - IP address
  - subnet mask
- Detects active interface automatically
- Tests DNS resolution
- Tests outbound internet connectivity
- Detects default gateway
- Fetches public IP
- Clean formatted CLI output
- Zero config
- Fast execution
- Cross-platform friendly (Go)

## Structure

- main.go            → CLI entrypoint
- pkg/network        → networking + connectivity logic
- pkg/utils          → system utilities
- binaries/          → compiled outputs
- recompile-all      → build helper script

## Build

go build -o netinfo

or

./recompile-all

## Run

./netinfo

## Output Includes

- Device name
- Interface list
- Active local IP
- Active interface
- Default gateway
- Public IP
- DNS status
- Internet connectivity status

## Use Cases

- Network diagnostics
- Host connectivity debugging
- Lab environments
- VM / container debugging
- Infra validation
- Local dev troubleshooting
- Automation pipelines
- Health checks
- CI environment validation

## Author

Christopher Buckner  
Cloud Infrastructure • Go • Networking • DevOps
