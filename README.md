# GO-Canary Token System
> **Advanced Deception & Forensic Monitoring Framework | Blue Team | SOC Tool**

![License: MIT](https://img.shields.io/badge/License-MIT-white.svg)
![Language: Go](https://img.shields.io/badge/Language-Go-black.svg)
![Security: Blue Team](https://img.shields.io/badge/Security-Blue_Team-grey.svg)

The **GO-Canary Token System** is a proactive defense framework designed to detect unauthorized internal reconnaissance. It utilizes "Honey-Assets" (digital tripwire documents) to lure intruders, capturing high-fidelity forensic metadata for real-time attribution and response.

---

## 🏗️ Architecture & Logic Flow
The system operates on a three-pillar security design to ensure stealth and reliability:

1.  **The Trap**: A deceptive `TOP_SECRET.html` file embedded with a hidden tracking pixel.
2.  **The Transit**: A secure **ngrok** tunnel that allows the "ping" to bypass local firewalls and NAT, enabling global tracking.
3.  **The Listener**: A high-performance Go-backend that extracts the attacker's **Public IP** (via `X-Forwarded-For`) and **User-Agent** fingerprint.
4.  **The Response**: Data is saved to an **SQLite** database while a background **Goroutine** dispatches an encrypted SMTP alert.

---

## 🛠️ Step 1: Prerequisites
Ensure the following are installed on your deployment machine:
* **Go (Golang)**: [Download here](https://go.dev/dl/)
* **ngrok**: [Download here](https://ngrok.com/download)
* **Gmail Account**: Required for receiving automated security alerts.

---

## ⚙️ Step 2: Security Configuration

### 1. Alerting Setup (SMTP)
Open `main.go` and locate the `const` section. Update these values:
```go
const (
    smtpEmail    = "your-email@gmail.com"     // Your Gmail address
    smtpPassword = "your-app-password"       // 16-character Google App Password
)
Security Warning: You must generate a "Google App Password" in your Google Account Security settings. Standard passwords will not be accepted.

2. Tunneling Setup (ngrok)
Launch the tunnel in your terminal: ngrok http 8080

Copy the Forwarding URL provided (e.g., https://a1b2.ngrok-free.app).

Open TOP_SECRET.html and update the fetch URL:

JavaScript
fetch("https://YOUR_NGROK_URL.ngrok-free.app/pixel.png", { ... });
Step 3: Deployment
Run the following commands in your project directory:

Windows (PowerShell):

PowerShell
$env:CGO_ENABLED = "0"
go mod tidy
go run main.go
Linux/macOS:

Bash
export CGO_ENABLED=0
go mod tidy
go run main.go
📊 Forensic Dashboard
Access the incident logs via the "Noir" forensic interface at: http://localhost:8080/view-logs.

Source IP: The attacker's true public origin.

Observed At: High-resolution forensic timestamps.

Fingerprint: Full browser and OS metadata for device attribution.

⚖️ Ethics & License
Distributed under the MIT License. For educational and authorized defensive security purposes only.


---

### **2. Step-by-Step Instructions for the User**
If you are explaining how to use this project to someone else, here is the simplified order:

1.  **Clone the Project**: Download the folder from GitHub.
2.  **Configure Credentials**:
    * The user must go to their Google Account and create an **App Password**.
    * Paste that email and password into the `const` section of `main.go`.
3.  **Set up the Tunnel**:
    * Open a terminal and run `ngrok http 8080`.
    * Take the link ngrok gives you and paste it into the `TOP_SECRET.html` file where it says `fetch`.
4.  **Run the Backend**:
    * Open a terminal in the folder.
    * Run `go run main.go`.
5.  **Trigger the Alert**:
    * Double-click `TOP_SECRET.html`.
    * The user will immediately receive an email alert, and the data will appear at `http://localhost:8080/view-logs`.

---

### **3. Final Upload Checklist**
Before you push this to your profile, make sure you have these **3 files** in your folder:
* **`.gitignore`**: Containing `alerts.db` (to keep your testing data private).
* **`LICENSE`**: Containing the MIT License text.
* **`README.md`**: The manual we just created.
