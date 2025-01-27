package mail

import (
	"errors"
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-notification-service/mail/model"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func (l *loginAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return "LOGIN", []byte{}, nil
}

func (l *loginAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(l.username), nil
		case "Password:":
			return []byte(l.password), nil
		default:
			return nil, errors.New("unexpected server challenge")
		}
	}
	return nil, nil
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

// SendMail is a function that sends an email.
func SendMail(payload model.MailPayload) error {
	host := config.GetString("SMTP_HOST", "smtp-mail.outlook.com")
	port := config.GetInt("SMTP_PORT", 587)
	username := config.GetString("SMTP_USERNAME", "annonymous@yopmail.com")
	password := config.GetString("SMTP_PASSWORD", "p@ssw0rd")
	address := fmt.Sprintf("%s:%d", host, port)

	to := []string{payload.To}
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		username,        // Sender's email address
		payload.To,      // Recipient's email address
		payload.Subject, // Subject of the email
		payload.Body,    // Body of the email
	))
	auth := LoginAuth(username, password)

	err := smtp.SendMail(address, auth, username, to, msg)
	if err != nil {
		return err
	}

	return nil
}
