package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"gopkg.in/gomail.v2"
)

func main() {
	body := "<!DOCTYPE html><html lang=\"en\"></html>"

	smtp_server := "smtp.exmail.qq.com"
	smtp_port := 465

	sender_email := "xxxxxx"
	sender_password := "xxxxxx"

	recipient_email := "xxxxxxxxx"

	message := gomail.NewMessage()
	message.SetHeader("From", sender_email)
	message.SetHeader("To", recipient_email)
	message.SetHeader("Subject", "test email")
	message.SetBody("text/html", body)

	d := gomail.NewDialer(smtp_server, smtp_port, sender_email, sender_password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
	fmt.Println("电子邮件已成功发送！")
}

type MailSender struct {
	Auth     smtp.Auth
	Sender   string
	SmtpAddr string
	InitFlag bool
}

func newSender() *MailSender {
	mailSender := new(MailSender)

	userEmail := "xxxxxx"
	mailSmtpPort := ":465"
	mailPassword := "xxxxxx"
	mailSmtpHost := "smtp.exmail.qq.com"

	// 这个是 对
	auth := smtp.PlainAuth("", userEmail, mailPassword, mailSmtpHost)

	sf := &smtp.ServerInfo{
		TLS:  true,
		Name: mailSmtpHost,
	}

	proto, bt, err := auth.Start(sf)
	if err != nil {
		log.Printf("smtp auth fail")
		return nil
	}
	log.Printf("proto = :%s", proto)
	log.Printf("bt =: %s", bt)

	mailSender.Auth = auth
	mailSender.SmtpAddr = fmt.Sprintf("%s%s", mailSmtpHost, mailSmtpPort)
	mailSender.Sender = userEmail
	mailSender.InitFlag = true
	return mailSender
}

func sendJiachaoEmail(to []string, body, nickname, subject string) {
	mailSender := newSender()
	if mailSender == nil || !mailSender.InitFlag {
	}
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + mailSender.Sender + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	log.Printf("start send email")

	err := smtp.SendMail(mailSender.SmtpAddr, mailSender.Auth, mailSender.Sender, to, msg)
	if err != nil {
		log.Fatalf("Send mail error: %v", err)
	}
	log.Printf("send success")
}
