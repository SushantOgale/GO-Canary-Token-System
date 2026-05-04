package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite" 
)

var db *sql.DB

const (
	// CONFIGURATION: Replace these before local execution
	smtpEmail    = "YOUR_EMAIL@gmail.com"     
	smtpPassword = "YOUR_APP_PASSWORD"       
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	dbPath       = "./alerts.db"
)

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("CRITICAL: Database connection failed:", err)
		return
	}
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

func sendProfessionalAlert(path, publicIP, localIP, ua string) {
	currentTime := time.Now().Format("02 Jan 2006, 15:04:05 MST")
	to := []string{smtpEmail}
	headerFrom := fmt.Sprintf("From: Security Monitor <%s>\r\n", smtpEmail)
	headerTo := fmt.Sprintf("To: %s\r\n", strings.Join(to, ", "))
	headerSubject := "Subject: [CRITICAL] Security Alert: Unauthorized Access!\r\n"
	headerMime := "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"
	body := fmt.Sprintf("Security Team,\n\nAn unauthorized access event recorded.\n\nForensic Details:\nPublic IP: %s\nLocal/Tunnel IP: %s\nAsset Path: %s\nTime: %s\nDevice: %s\n\nAction: Review logs at http://localhost:8080/view-logs", publicIP, localIP, path, currentTime, ua)
	
	message := []byte(headerFrom + headerTo + headerSubject + headerMime + body)
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, to, message)
	if err != nil {
		fmt.Println("SMTP Error:", err)
	} else {
		fmt.Printf("[%s] Alert dispatched: %s\n", currentTime, path)
	}
}

func logToDB(path, publicIP, localIP, ua string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO incidents (timestamp, path, ip, proxy_ip, useragent) VALUES (?, ?, ?, ?, ?)", timestamp, path, publicIP, localIP, ua)
	if err != nil {
		fmt.Println("DB Error:", err)
	}
	go sendProfessionalAlert(path, publicIP, localIP, ua)
}

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
<html>
<head>
    <title>SECURITY | ALERTS</title>
    <style>
        :root { --bg: #000000; --surface: #111111; --border: #333333; --text: #ffffff; --dim: #888888; }
        body { margin: 0; padding: 40px; font-family: 'Segoe UI', sans-serif; background: var(--bg); color: var(--text); display: flex; flex-direction: column; align-items: center; }
        h1 { letter-spacing: 10px; text-transform: uppercase; border-bottom: 1px solid var(--border); padding-bottom: 20px; width: 100%%; max-width: 1000px; text-align: center; }
        table { width: 100%%; max-width: 1000px; border-collapse: collapse; background: var(--surface); border: 1px solid var(--border); }
        th { padding: 15px; text-align: left; font-size: 10px; color: var(--dim); text-transform: uppercase; border-bottom: 1px solid var(--border); }
        td { padding: 20px; border-bottom: 1px solid var(--border); font-size: 13px; }
        .mono { font-family: monospace; font-weight: bold; }
        .badge { border: 1px solid white; padding: 2px 8px; font-size: 10px; }
    </style>
</head>
<body>
    <h1>ALERTS</h1>
    <table>
        <thead><tr><th>Identifier</th><th>Source IP</th><th>Observed At</th><th>Fingerprint</th></tr></thead>
        <tbody>`)

	for rows.Next() {
		var ts, path, pub, loc, ua string
		rows.Scan(&ts, &path, &pub, &loc, &ua)
		fmt.Fprintf(w, `<tr><td><span class="badge">%s</span></td><td class="mono">%s</td><td class="mono">%s</td><td style="color:#888">%s</td></tr>`, path, pub, ts, ua)
	}
	fmt.Fprintf(w, `</tbody></table></body></html>`)
}

func requestInterceptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.URL.Path == "/view-logs" {
		viewLogsHandler(w, r)
		return
	}
	localIP := r.RemoteAddr
	publicIP := r.Header.Get("X-Forwarded-For")
	if publicIP == "" { publicIP = localIP }
	logToDB(r.URL.Path, publicIP, localIP, r.UserAgent())
	if strings.HasSuffix(r.URL.Path, ".png") {
		w.Header().Set("Content-Type", "image/png")
		pixel, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAACIHAWCAAAAC0IEQVR42mNkYAAAAAYAAJCB0C8AAAAASUVORK5CYII=")
		w.Write(pixel)
	}
}

func main() {
	initDatabase()
	http.HandleFunc("/", requestInterceptor)
	fmt.Println("SYSTEM ONLINE | DASHBOARD: http://localhost:8080/view-logs")
	http.ListenAndServe(":8080", nil)
}
