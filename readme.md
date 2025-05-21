
# Renovate Playground for Antrea – LFX Mentorship #7155

A minimal Go program demonstrating Renovate’s dependency management using a vulnerable version of `miekg/dns`, aligned with Antrea's networking focus.

## 🔍 What It Does
- **DNS Resolution**: Resolves `example.com` using `miekg/dns`.
- **HTTP Check**: Verifies accessibility via HTTP GET.
- **Output**: Prints `UP` or `DOWN` with detailed reasons.

## ⚠️ Vulnerable Dependency
- **Module**: `github.com/miekg/dns`
- **Version**: `v1.1.43`
- **CVE**: [CVE-2021-36659](https://github.com/advisories/GHSA-4jv9-86wv-5v3h)
- **Fixed In**: `v1.1.50+`

## 🚀 Renovate Features
- Configured via `renovate.json`
- Detects CVEs, raises security PRs
- Generates a Dependency Dashboard issue
- Integrated with Mend Developer Portal (v40.16.0)

## 🧪 How to Run
```bash
git clone https://github.com/gitsofaryan/renovate-playground.git
cd renovate-playground
go mod tidy
go run main.go
````

## ✅ Example Output

```text
DNS resolved example.com to 93.184.216.34
Status for example.com: UP (HTTP 200)
```

## 📁 Project Structure

```
renovate-playground/
├── main.go
├── go.mod
├── go.sum
├── renovate.json
└── README.md
```

## 🛠️ Setup Renovate

1. Install [Renovate GitHub App](https://github.com/apps/renovate)
2. Enable it for `renovate-playground`
3. Monitor PRs and issues
4. Track jobs on [Mend Developer Portal](https://developer.mend.io/)







