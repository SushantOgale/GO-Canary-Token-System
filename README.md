GO-Canary Token System
Advanced Deception & Forensic Monitoring Framework | BLUE TEAM | SOC TOOL

The GO-Canary Token System is a proactive defense framework designed to detect unauthorized internal reconnaissance. It utilizes "Honey-Assets" (digital tripwire documents) to lure intruders, capturing high-fidelity forensic metadata for real-time attribution.

Architecture Overview

1) The system follows a three-pillar security design:

2) The Trap: Deceptive assets (HTML/Honeyfiles) embedded with tracking pixels.

3) The Tunnel: Secure ngrok relay for bypassing firewalls and NAT.

4) The Listener: High-performance Go-backend for forensic extraction and SMTP alerting.

🛠️ Step 1: Prerequisites
Before deployment, ensure the following are installed:

1) Go (Golang): https://go.dev/dl/ 

2) ngrok: https://ngrok.com/download  

- Gmail Account: Required for receiving automated security alerts.

⚙️ Step 2: Security Configuration (Important)
1. Alerting Setup (SMTP)
Open main.go and locate the const section. Update these values:

Go
const (
    smtpEmail    = "your-email@gmail.com"     // Your Gmail address
    smtpPassword = "your-app-password"       // Your 16-character Google App Password
)
⚠️ Security Warning: You must generate a "Google App Password" in your Google Account Security settings. Your standard Gmail password will not work.

2. Tunneling Setup (ngrok)
Open a new terminal and run: ngrok http 8080

Copy the Forwarding URL provided (e.g., https://a1b2-c3d4.ngrok-free.app).

Open TOP_SECRET.html and replace the placeholder URL:
fetch("https://YOUR_NGROK_URL_HERE/pixel.png", { ... });

🚀 Step 3: Deployment
Initialize Environment:

Bash
$env:CGO_ENABLED = "0"
go mod tidy
Start the Listener:

Bash
go run main.go
📊 Forensic Dashboard
Access the incident logs and "Noir" interface at: http://localhost:8080/view-logs.

1) Identifier: The specific trap triggered.

2) Source IP: Attacker's true public IP (via X-Forwarded-For).

3) Observed At: High-resolution timestamp.

4) Fingerprint: Full browser/OS User-Agent for attribution.

⚖️ License
This project is licensed under the MIT License - see the LICENSE file for details.

2. Essential Repository Tips
The .gitignore File: You must have this file to ensure your private database (alerts.db) and email logs aren't uploaded.

The LICENSE File: Professional repos always include a license. I suggest the MIT License—it’s short, simple, and standard.

Repository Tags: On the right side of your GitHub repo, add "Topics":

1) cybersecurity

2) forensics

3) golang

4) intrusion-detection

5) honeypot
