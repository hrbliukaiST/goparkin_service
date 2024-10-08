package email

import (
	"goparkin_service/config"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", config.AppConfig.EmailUser)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, config.AppConfig.EmailUser, config.AppConfig.EmailPass)

    return d.DialAndSend(m)
}