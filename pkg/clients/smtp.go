package clients

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type MailInfo struct {
	AlertName   string
	Projectid   string
	Description string
}

type SmtpClient struct {
	Host string
	auth sasl.Client
}

func NewSmtpClient(host string, username string, secret string) *SmtpClient {
	auth := sasl.NewPlainClient("", username, secret)
	return &SmtpClient{
		Host: host,
		auth: auth,
	}
}

func (sc *SmtpClient) SendEmail(from string, to []string, info *MailInfo) error {
	// TODO: Fix path if needed (docker)
	wd, _ := os.Getwd()
	tmpl, err := template.ParseFiles(wd + "/../clients/templates/template.html")
	if err != nil {
		logger.Printf("failed to read email template")
		return err
	}

	var html bytes.Buffer
	tmpl.Execute(&html, struct {
		AlertName   string
		Project     string
		Description string
	}{
		AlertName:   info.AlertName,
		Project:     info.Projectid,
		Description: info.Description,
	})

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	htmlstring := html.String()

	var msg bytes.Buffer
	fmt.Fprintf(&msg,
		"Subject: Alert triggered: %s for project:%s \r\n"+
			"From: %s\r\n"+
			"To: %s\r\n"+
			"%s\r\n"+
			"%s\r\n",
		info.AlertName, info.Projectid, from, to[0], headers, htmlstring)

	reader := bytes.NewReader(msg.Bytes())
	msgstring := msg.String()
	logger.Printf(msgstring)
	err = smtp.SendMail(sc.Host, sc.auth, from, to, reader)
	if err != nil {
		logger.Printf("failed to send email")
		return err
	}
	logger.Printf("email has been sent successfully")
	return nil
}
