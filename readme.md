# Renovate Playground for Antrea #7155

A minimal Go program demonstrating Renovate’s dependency management for the Antrea 2025 LFX Mentorship program (#7155).

## What It Does
The program (`main.go`) monitors the status of “example.com” by checking DNS resolution and HTTP reachability using `github.com/miekg/dns@v1.1.43`.

## How to Run
1. Ensure Go 1.21+ is installed and internet access is available.
2. Run the program:
   ```bash
   go run main.go

## Output (example):
```bash
DNS resolved example.com to 93.184.216.34
Status for example.com: UP (HTTP 200)
```