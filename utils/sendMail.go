package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"strconv"

	"github.com/guatom999/TicketShop-Movie/config"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
}

func SendEmail(cfg *config.Config, to string, subject string, body string) error {
	from := cfg.Mailer.MailerUserName
	password := cfg.Mailer.MailerPassword

	smtpHost := cfg.Mailer.MailerHost
	smtpPort := cfg.Mailer.MailerPort

	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func SecondSendEmail(cfg *config.Config, message *gomail.Message) error {

	mailer := ConnectToMailer(cfg)

	message.SetHeader("From", "bossbadz642@gmail.com")

	if err := mailer.DialAndSend(message); err != nil {
		log.Panicln("[Mailer] ", err)
		return errors.New("failed to send email")
	}

	return nil

}

func ConnectToMailer(cfg *config.Config) *gomail.Dialer {
	mailer := gomail.NewDialer(cfg.Mailer.MailerHost, cfg.Mailer.MailerPort, cfg.Mailer.MailerUserName, cfg.Mailer.MailerPassword)

	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return mailer

}
