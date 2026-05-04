
```markdown
# GO-Canary Token System
Advanced Deception & Forensic Monitoring Framework | Blue Team | SOC Tool

![License: MIT](https://img.shields.io/badge/License-MIT-white.svg)
![Language: Go](https://img.shields.io/badge/Language-Go-black.svg)
![Security: Blue Team](https://img.shields.io/badge/Security-Blue_Team-grey.svg)

The GO-Canary Token System is a proactive defense framework designed to detect unauthorized internal reconnaissance. It utilizes "Honey-Assets" to lure intruders, capturing high-fidelity forensic metadata for real-time attribution and response.

---

🏗️ Architecture & Logic Flow
The system operates on a three-pillar security design:

* **The Trap**: A deceptive `TOP_SECRET.html` file embedded with a hidden tracking pixel.
* **The Transit**: A secure **ngrok** tunnel that bypasses local firewalls and NAT.
* **The Listener**: A high-performance Go-backend that extracts **Public IP** and **User-Agent**.
* **The Response**: Data is saved to **SQLite** while a **Goroutine** dispatches an SMTP alert.

---

🛠️ Step 1: Prerequisites
Ensure the following are installed on your deployment machine:
1. **Go (Golang)**: [Download here](https://go.dev/dl/)
2. **ngrok**: [Download here](https://ngrok.com/download)
3. **Gmail Account**: Required for receiving automated security alerts.

---

⚙️ Step 2: Configuration

1. Alerting Setup (SMTP)
Open `main.go` and update your credentials in the `const` section:
```go
const (
    smtpEmail    = "your-email@gmail.com"
    smtpPassword = "your-app-password"
)
```
> Security Note:** You must generate a "Google App Password" in your Google Account settings. Standard passwords will not work.

2. Tunneling Setup (ngrok)
1. Launch Tunnel: Run `ngrok http 8080` in your terminal.
2. Capture URL: Copy the Forwarding URL (e.g., `https://a1b2.ngrok-free.app`).
3. Update Asset: Open `TOP_SECRET.html` and replace the placeholder:
   `fetch("https://YOUR_NGROK_URL_HERE/pixel.png", { ... });`

---

Step 3: Deployment Commands
Run these commands in your terminal to launch the system:

Windows (PowerShell)
```powershell
$env:CGO_ENABLED = "0"
go mod tidy
go run main.go
```

Linux / macOS (Bash)
```bash
export CGO_ENABLED=0
go mod tidy
go run main.go
```

---

Forensic Dashboard
Once live, access the incident logs via the "Noir" interface:
URL: `http://localhost:8080/view-logs`

Captured Metadata:
* Source IP: The attacker's true public origin (bypassing proxies).
* Observed At: High-resolution forensic timestamps.
* Fingerprint: Full browser and OS metadata for device attribution.

---

## ⚖️ License & Ethics
Distributed under the **MIT License**. 

Disclaimer: This tool is for educational and authorized defensive security purposes only. The author is not responsible for any misuse. Always ensure you have explicit permission before deploying deception assets.
```


Final Polish for your GitHub Profile:

1.  Add a "Social Preview": Go to your Repo Settings > General > Social Preview. Upload a screenshot of your Black ALERTS Dashboard. It makes the repo look amazing when shared on LinkedIn.
2.  Add Topics: On your main repo page, click the gear icon next to "About" and add these exactly: `cybersecurity`, `blue-team`, `honeypot`, `golang`, `forensics`.
3.  Delete `alerts.db`: I noticed your repo has `alerts.db` uploaded. Delete it manually from GitHub.Then, ensure your `.gitignore` file has the word `alerts.db` inside it. A professional repo should never contain the developer's personal test logs!

Does this look like the "Perfect" version you were aiming for?
