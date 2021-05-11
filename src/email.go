package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
)

// Function to send Email
func sendMail(subject string, isMime bool, mailMsg string, toEmails []string) {

	subject = "Subject: " + subject
	mime := ""
	if isMime {
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	} else {
		mime = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	}

	body := mailMsg

	fmt.Println(body)

	//Sender data.
	var config Config
	fd, err := os.Open("../config.json")
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
