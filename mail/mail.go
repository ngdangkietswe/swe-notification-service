package mail

import (
	"bytes"
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

	var emailBuf bytes.Buffer
	emailBuf.WriteString(fmt.Sprintf("To: %s\r\n", payload.To))
	emailBuf.WriteString(fmt.Sprintf("From: %s\r\n", username)) // Adjust sender
	emailBuf.WriteString(fmt.Sprintf("Subject: %s\r\n", payload.Subject))
	emailBuf.WriteString("MIME-Version: 1.0\r\n")
	emailBuf.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	emailBuf.WriteString("\r\n") // Blank line separates headers from body
	emailBuf.WriteString(payload.Body)

	to := []string{payload.To}
	auth := LoginAuth(username, password)

	err := smtp.SendMail(address, auth, username, to, emailBuf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
