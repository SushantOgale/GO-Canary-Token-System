package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite" // CGO-free driver for Windows portability
)

var db *sql.DB

const (
	smtpEmail    = "fakemailer03@gmail.com"
	smtpPassword = "xeqsxgbjdzswfyra"
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	dbPath       = "./alerts.db"
)

// Orchestrator: System Initialization
func initDatabase() {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("CRITICAL: Database connection failed:", err)
		return
	}
	// Verify and create schema
	query := `CREATE TABLE IF NOT EXISTS incidents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TEXT,
		path TEXT,
		ip TEXT,
		proxy_ip TEXT,
		useragent TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("CRITICAL: Table creation failed:", err)
	}
}

// Alerting: Asynchronous Notification
func sendProfessionalAlert(path, publicIP, localIP, ua string) {
	currentTime := time.Now().Format("02 Jan 2006, 15:04:05 MST")
	to := []string{smtpEmail}
	headerFrom := fmt.Sprintf("From: Security Monitor <%s>\r\n", smtpEmail)
	headerTo := fmt.Sprintf("To: %s\r\n", strings.Join(to, ", "))
	headerSubject := "Subject: [CRITICAL] Security Alert: Unauthorized Access!\r\n"
	headerMime := "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"
	body := fmt.Sprintf("Security Team,\n\nAn unauthorized access event has been recorded.\n\nForensic Details:\nPublic/Attacker IP: %s\nLocal/Tunnel IP: %s\nAsset Path: %s\nTime: %s\nDevice: %s\n\nAction: Review logs at http://localhost:8080/view-logs", publicIP, localIP, path, currentTime, ua)

	message := []byte(headerFrom + headerTo + headerSubject + headerMime + body)
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, to, message)
	if err != nil {
		fmt.Println("SMTP Error:", err)
	} else {
		fmt.Printf("[%s] Alert sent for path: %s\n", currentTime, path)
	}
}

// Persistence: Forensic Evidence Archival
func logToDB(path, publicIP, localIP, ua string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO incidents (timestamp, path, ip, proxy_ip, useragent) VALUES (?, ?, ?, ?, ?)", timestamp, path, publicIP, localIP, ua)
	if err != nil {
		fmt.Println("DB Error:", err)
	}
	// Use Goroutine for non-blocking email alert
	go sendProfessionalAlert(path, publicIP, localIP, ua)
}

// Dashboard: Black and White Forensic Interface
func viewLogsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT timestamp, path, ip, proxy_ip, useragent FROM incidents ORDER BY id DESC")
	if err != nil {
		http.Error(w, "Database Error", 500)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SECURITY | ALERTS</title>
    <style>
        :root {
            --bg: #000000;
            --surface: #111111;
            --border: #333333;
            --text-main: #ffffff;
            --text-dim: #888888;
        }
        body {
            margin: 0;
            padding: 60px 20px;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: var(--bg);
            color: var(--text-main);
            display: flex;
            flex-direction: column;
            align-items: center;
        }
        .title-container {
            text-align: center;
            margin-bottom: 50px;
            border-bottom: 1px solid var(--border);
            padding-bottom: 20px;
            width: 100%%;
            max-width: 1200px;
        }
        h1 {
            font-size: 2.5rem;
            font-weight: 200;
            letter-spacing: 12px;
            text-transform: uppercase;
            margin: 0;
        }
        .main-table {
            width: 100%%;
            max-width: 1200px;
            background: var(--surface);
            border: 1px solid var(--border);
            border-radius: 4px;
            border-collapse: collapse;
        }
        th {
            background: #000000;
            padding: 20px;
            text-align: left;
            font-size: 11px;
            color: var(--text-dim);
            text-transform: uppercase;
            letter-spacing: 2px;
            border-bottom: 1px solid var(--border);
        }
        td {
            padding: 25px 20px;
            border-bottom: 1px solid var(--border);
            font-size: 14px;
            color: var(--text-main);
        }
        tr:hover td {
            background: #181818;
        }
        .mono {
            font-family: 'Courier New', Courier, monospace;
            font-weight: bold;
        }
        .badge {
            border: 1px solid var(--text-main);
            padding: 4px 10px;
            font-size: 11px;
            font-weight: bold;
            text-transform: uppercase;
        }
        .ua-text {
            color: var(--text-dim);
            font-size: 12px;
            max-width: 450px;
            line-height: 1.4;
        }
    </style>
</head>
<body>
    <div class="title-container">
        <h1>ALERTS</h1>
    </div>

    <table class="main-table">
        <thead>
            <tr>
                <th>Identifier</th>
                <th>Source IP & Metadata</th>
                <th>Observed At</th>
                <th>Client Fingerprint</th>
            </tr>
        </thead>
        <tbody>`)

	for rows.Next() {
		var ts, path, pub, loc, ua string
		rows.Scan(&ts, &path, &pub, &loc, &ua)
		fmt.Fprintf(w, `
            <tr>
                <td><span class="badge">%s</span></td>
                <td>
                    <div class="mono">%s</div>
                    <div style="font-size: 10px; color: #555;">RELAY: %s</div>
                </td>
                <td class="mono">%s</td>
                <td class="ua-text">%s</td>
            </tr>`, path, pub, loc, ts, ua)
	}

	fmt.Fprintf(w, `
        </tbody>
    </table>
</body>
</html>`)
}

// Interceptor: Forensic Extraction & Payload Delivery
func requestInterceptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Routing for the Dashboard
	if r.URL.Path == "/view-logs" {
		viewLogsHandler(w, r)
		return
	}

	// Forensic Data Collection: Capture True Public IP
	localIP := r.RemoteAddr
	publicIP := r.Header.Get("X-Forwarded-For")
	if publicIP == "" {
		publicIP = localIP
	}

	// Log incident and trigger background alert
	logToDB(r.URL.Path, publicIP, localIP, r.UserAgent())

	// Deliver Stealth Pixel Payload
	if strings.HasSuffix(r.URL.Path, ".png") {
		w.Header().Set("Content-Type", "image/png")
		// Invisible 1x1 Transparent PNG
		pixel, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAACIHAWCAAAAC0IEQVR42mNkYAAAAAYAAJCB0C8AAAAASUVORK5CYII=")
		w.Write(pixel)
		return
	}
}

func main() {
	initDatabase()
	http.HandleFunc("/", requestInterceptor)

	fmt.Println("------------------------------------")
	fmt.Println("  SECURITY MONITOR: ONLINE")
	fmt.Println("  Dashboard URL: http://localhost:8080/view-logs")
	fmt.Println("------------------------------------")

	http.ListenAndServe(":8080", nil)
}
