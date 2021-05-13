package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Config struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	RecaptchaSecret string `json:"recaptcha_secret"`
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

var database *sql.DB

func main() {
	database, err := sql.Open("sqlite3", "getcowinslot/data/db/user_data.db")
	if err != nil {
		fmt.Println("Database Error: %v", err)
	}
	defer database.Close()
	initDB()
}

func initDB() {
	var email string

	rows, err := database.Query("SELECT DITINCT emails FROM users")
	if err != nil {
		fmt.Println("Statement InitDB Error: ", err)
		return
	}

	for rows.Next() {
		rows.Scan(&email)
		emailArr := strings.Split(email, ",")
		go sendMail(emailArr)
		time.Sleep(120 * time.Second)
	}

}

// Function to send Email
func sendMail(toEmails []string) {

	subject := "Subject: Email notifications disabled\n"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	mailMsg := "Hi,<br><br>We tried our best to provide you real time vaccine availabilty notifications. But it has become almost impossible for us, as the public APIs are blocking our bots for making too many requests. This might be our last email.<br><br><b>However, we have released a browser only functionality: https://getcowinslot.in/client-notify.html</b><br>Keep a tab open in your browser and it will alert you whenever a slot opens up.<br><br>Stay aware of the availabilty, and get yourself vaccinated as soon as possible. 🙂<br><br>-Hrushikesh<br><br>Source: https://www.indiatoday.in/technology/news/story/changes-in-cowin-app-govt-restricts-vaccine-slot-info-to-fight-bots-and-alert-services-1799827-2021-05-07<br>"

	body := mailMsg

	fmt.Println(body)

	//Sender data.
	var config Config
	fd, err := os.Open("getcowinslot/config.json")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	creds, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err)
	}

	credsString := fmt.Sprintf("%s", creds)

	err = json.Unmarshal([]byte(credsString), &config)
	from := config.Email
	password := config.Password

	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte(subject + mime + body)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err = smtp.SendMail(smtpServer.Address(), auth, from, toEmails, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent to ", toEmails)

}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
