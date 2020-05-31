package handlers

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func sendEmail(subject string, from string, data map[string]string) (bool, error) {
	host := os.Getenv("API_SMTP_HOST")
	user := os.Getenv("API_SMTP_USER")
	password := os.Getenv("API_SMTP_PASSWORD")
	sender := os.Getenv(("API_SMTP_SENDER_ADDRESS"))
	port, err := strconv.Atoi(os.Getenv("API_SMTP_PORT"))
	if err != nil {
		return false, err
	}

	// message body
	body := "From: " + from
	for key, val := range data {
		body = fmt.Sprintf("%s\n\n%s:\n%s", body, key, val)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", sender)
	m.SetHeader("Subject", fmt.Sprintf("[shareaudio.cc] %s", subject))
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(host, port, user, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}
