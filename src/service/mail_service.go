package service

import (
	gomail "gopkg.in/mail.v2"
	"twodo.app/condo/model"
)

type mailService struct {
	SendMail func(attachment string, body string, subject string) error
}

var mail = mailService{
	SendMail: func(attachment string, body string, subject string) error {
		config := model.LoadConfig()
		smtpConfig := config.Smtp

		m := gomail.NewMessage()
		m.Attach(attachment)
		m.SetHeader("From", smtpConfig.From)
		m.SetHeader("To", config.CondoEmail)
		m.SetHeader("Bcc", smtpConfig.From)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)

		d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Username, smtpConfig.Password)

		return d.DialAndSend(m)
	},
}
