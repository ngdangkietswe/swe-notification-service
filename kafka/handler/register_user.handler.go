package handler

import (
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"log"
)

type RegisterUserHandler struct {
}

func NewRegisterUserHandler() *RegisterUserHandler {
	return &RegisterUserHandler{}
}

// Handle sends a welcome email to the user.
func (h *RegisterUserHandler) Handle(payload *domain.RegisterUser) {
	log.Printf("[REGISTER USER HANDLER] Processing registration for email: %s", payload.Email)

	// Prepare the email payload.
	emailPayload := mailpayload.MailPayload{
		To:      payload.Email,
		Subject: "Welcome to SWE Notification Service",
		Body: fmt.Sprintf(
			"Hello %s,\n\nWelcome to SWE!\n\nYour account has been successfully created at %s with the username: %s.\n\nBest regards,\nSWE Notification Service Team",
			payload.Username,
			payload.CreatedAt,
			payload.Username,
		),
	}

	log.Printf("[REGISTER USER HANDLER] Sending welcome email to %s", payload.Email)
	mail.SendMail(emailPayload)
}
