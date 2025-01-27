package handler

import (
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"go.uber.org/zap"
)

type RegisterUserHandler struct {
	logger *logger.Logger
}

func NewRegisterUserHandler(logger *logger.Logger) *RegisterUserHandler {
	return &RegisterUserHandler{
		logger: logger,
	}
}

// Handle sends a welcome email to the user.
func (h *RegisterUserHandler) Handle(payload *domain.RegisterUser) {
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

	h.logger.Info("Sending welcome email to user", zap.String("email", payload.Email))

	if err := mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send welcome email to user", zap.String("email", payload.Email), zap.Error(err))
	}
}
