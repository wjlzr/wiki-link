package email

import (
	"crypto/tls"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var SmtpConfig Smtp

type Smtp struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ContentType string `json:"content_type"`
}

func SendMail(to string, Subject, bodyMessage string) (err error) {
	d := gomail.NewDialer(
		SmtpConfig.Host,
		SmtpConfig.Port,
		SmtpConfig.Username,
		SmtpConfig.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", SmtpConfig.Username)
	m.SetAddressHeader("From", SmtpConfig.Username, "Wiki Link Team")
	m.SetHeader("To", to)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", bodyMessage)
	if err = gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %q: %v", to, err)
	}
	return
}

func InitSmtp(cfg *viper.Viper) {
	SmtpConfig = Smtp{
		Host:        cfg.GetString("host"),
		Port:        cfg.GetInt("port"),
		Username:    cfg.GetString("username"),
		Password:    cfg.GetString("password"),
		ContentType: cfg.GetString("content_type"),
	}
}
