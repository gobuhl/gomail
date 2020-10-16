package gomail

import (
	"errors"
	"fmt"
	"net/smtp"
)

var emailAuth smtp.Auth

type EmailData struct {
	from     string
	host     string
	password string
	port     string
}

func SendEmailSMTP(to []string, data interface{}, templatePath string, emailData EmailData) (bool, error) {
	emailHost := emailData.host
	emailFrom := emailData.from
	emailPassword := emailData.password
	emailPort := emailData.port

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(templatePath, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}
