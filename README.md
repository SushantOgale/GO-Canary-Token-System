# GO-Canary Token System
**Advanced Deception & Forensic Monitoring Framework | Blue Team | SOC Tool**

![License: MIT](https://img.shields.io/badge/License-MIT-white.svg)
![Language: Go](https://img.shields.io/badge/Language-Go-black.svg)
![Security: Blue Team](https://img.shields.io/badge/Security-Blue_Team-grey.svg)

The **GO-Canary Token System** is a proactive defense framework designed to detect unauthorized internal reconnaissance. It utilizes "Honey-Assets" (digital tripwire documents) to lure intruders, capturing high-fidelity forensic metadata for real-time attribution and response.

---

## Architecture & Logic Flow
The system operates on a three-pillar security design:

1. **The Trap**: A deceptive `TOP_SECRET.html` file embedded with a hidden tracking pixel.
2. **The Transit**: A secure **ngrok** tunnel that allows the "ping" to bypass local firewalls and NAT.
3. **The Listener**: A high-performance Go-backend that extracts the attacker's **Public IP** (via X-Forwarded-For) and **User-Agent** fingerprint.
4. **The Response**: Data is saved to an **SQLite** database while a background **Goroutine** dispatches an encrypted SMTP alert to the SOC administrator.

---

## Step 1: Prerequisites
Ensure the following are installed on your deployment machine:
1. **Go (Golang)**: [Download here](https://go.dev/dl/)
2. **ngrok**: [Download here](https://ngrok.com/download)
3. **Gmail Account**: Required for receiving automated security alerts.

---

## Step 2: Configuration (Mandatory Changes)

### 1. Alerting Setup (SMTP)
Open `main.go` and locate the `const` section. Update these values with your credentials:

```go
const (
    smtpEmail    = "your-email@gmail.com"     // Your Gmail address
    smtpPassword = "your-app-password"       // 16-character Google App Password
)

Security Note: You must generate a "Google App Password" in your Google Account Security settings. Standard Gmail passwords will not work.

2. Tunneling Setup (ngrok)
Open a terminal and launch the tunnel:
ngrok http 8080

Copy the Forwarding URL provided (e.g., https://a1b2-c3d4.ngrok-free.app).

Open TOP_SECRET.html and replace the placeholder URL in the script:
fetch("https://YOUR_NGROK_URL_HERE/pixel.png", { ... });
Step 3: Deployment (Terminal Commands)
Copy and paste these commands into your terminal to launch the system.

For Windows (PowerShell):
# Set environment for portability
$env:CGO_ENABLED = "0"

# Install dependencies
go mod tidy

# Launch the listener
go run main.go

For Linux/macOS (Bash):

Bash
# Set environment for portability
export CGO_ENABLED=0

# Install dependencies
go mod tidy

# Launch the listener
go run main.go


Forensic Dashboard
Once the system is live, access the incident logs via the forensic interface at:
http://localhost:8080/view-logs

Data Captured:

Identifier: The specific asset path triggered.

Source IP: The attacker's true public origin (bypassing proxies).

Observed At: High-resolution forensic timestamps.

Fingerprint: Full browser and OS metadata for device attribution.

License & Ethics
Distributed under the MIT License.

Disclaimer: This tool is developed for educational and authorized defensive security purposes only. The author is not responsible for any misuse or damage caused by this software. Always ensure you have explicit permission before deploying deception assets on a network.

