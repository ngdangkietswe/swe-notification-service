package handler

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/constants"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
	"github.com/ngdangkietswe/swe-notification-service/data/repository/emailtemplate"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"strings"
)

type ResetPasswordHandler struct {
	logger            *logger.Logger
	emailTemplateRepo emailtemplate.IEmailTemplateRepository
}

func NewResetPasswordHandler(logger *logger.Logger, emailTemplateRepo emailtemplate.IEmailTemplateRepository) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		logger:            logger,
		emailTemplateRepo: emailTemplateRepo,
	}
}

// Handle sends a reset password email to the user.
func (h *ResetPasswordHandler) Handle(payload *domain.ResetPassword) {
	emailTemplate, err := h.emailTemplateRepo.FindByKey(context.Background(), constants.EmailTemplateKeyResetPassword)
	if err != nil {
		h.logger.Error(
			lo.Ternary(ent.IsNotFound(err), "Email template not found", "Failed to get Email template"),
			zap.String("error", err.Error()),
		)
		return
	}

	// Prepare the email payload.
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", payload.Token)

	body := strings.NewReplacer(
		"{reset_link}", resetLink,
	).Replace(emailTemplate.Body)

	emailPayload := mailpayload.MailPayload{
		To:      payload.Email,
		Subject: emailTemplate.Subject,
		Body:    body,
	}

	h.logger.Info("Sending email for reset password", zap.String("email", payload.Email))

	if err = mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send email for reset password", zap.String("email", payload.Email), zap.Error(err))
		return
	}

	h.logger.Info("Reset password email sent successfully", zap.String("email", payload.Email))
}
