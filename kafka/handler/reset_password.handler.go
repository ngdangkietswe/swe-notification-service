package handler

import (
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"log"
)

type ResetPasswordHandler struct {
}

func NewResetPasswordHandler() *ResetPasswordHandler {
	return &ResetPasswordHandler{}
}

// Handle sends a reset password email to the user.
func (h *ResetPasswordHandler) Handle(payload *domain.ResetPassword) {
	log.Printf("[RESET PASSWORD HANDLER] Processing reset password for email: %s", payload.Email)

	// Prepare the email payload.
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", payload.Token)
	emailPayload := mailpayload.MailPayload{
		To:      payload.Email,
		Subject: "Reset Password Request",
		Body: fmt.Sprintf(
			"Hello,\n\nWe received a request to reset your password. If you did not make this request, please ignore this email.\n\nTo reset your password, click the link below:\n%s\n\nIf you're having trouble clicking the link, copy and paste the URL into your browser.\n\nThis link will expire in 30 minutes.\n\nBest regards,\nSWE Notification Service Team",
			resetLink,
		),
	}

	log.Printf("[RESET PASSWORD HANDLER] Sending reset password to %s", payload.Email)
	mail.SendMail(emailPayload)
}
