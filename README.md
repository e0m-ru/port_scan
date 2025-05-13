# Port Scanner - Simple TCP Port Scanner

## Overview
A lightweight TCP port scanner written in Go. This tool allows you to quickly check open ports on a target host within a specified range.

## Features
- Scans TCP ports on a given host
- Customizable port range (single port or range)
- Adjustable timeout for connection attempts
- Concurrent scanning for faster results
- Simple command-line interface
- Clear output showing open/closed ports

## Usage
```bash
./tcp_scaner [-p <port-range> -t <miliseconds> -w <GorutinsCount>] <target>
```
## Examples:
```bash
# Scan ports 1-100 on localhost
./tcp_scaner -p 1-100 localhost
```
```bash
# Scan single port with custom timeout
./tcp_scaner  -p 80 -timeout 200 example.com
```
```bash
# Scan specific ports
./tcp_scaner -p 22,80,443 192.168.1.1
```

## Installation
1. Clone the repository:
```bash
git clone https://github.com/e0m/port_scan.git
```
2. Build the project:
```bash
cd port_scan
go build -o tcp_scaner
```


## Dependencies
Go 1.16 or higher

## License
MIT License - see [LICENSE](https://license/) for details
