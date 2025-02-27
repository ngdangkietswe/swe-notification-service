package handler

import (
	"context"
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

type RegisterUserHandler struct {
	logger            *logger.Logger
	emailTemplateRepo emailtemplate.IEmailTemplateRepository
}

func NewRegisterUserHandler(logger *logger.Logger, emailTemplateRepo emailtemplate.IEmailTemplateRepository) *RegisterUserHandler {
	return &RegisterUserHandler{
		logger:            logger,
		emailTemplateRepo: emailTemplateRepo,
	}
}

// Handle sends a welcome email to the user.
func (h *RegisterUserHandler) Handle(payload *domain.RegisterUser) {
	emailTemplate, err := h.emailTemplateRepo.FindByKey(context.Background(), constants.EmailTemplateKeyRegisterUser)
	if err != nil {
		h.logger.Error(
			lo.Ternary(ent.IsNotFound(err), "Email template not found", "Failed to get Email template"),
			zap.String("error", err.Error()),
		)
		return
	}

	body := strings.NewReplacer(
		"{username}", payload.Username,
		"{created_at}", payload.CreatedAt,
	).Replace(emailTemplate.Body)

	// Prepare the email payload.
	emailPayload := mailpayload.MailPayload{
		To:      payload.Email,
		Subject: emailTemplate.Subject,
		Body:    body,
	}

	h.logger.Info("Sending welcome email to user", zap.String("email", payload.Email))

	if err = mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send welcome email to user", zap.String("email", payload.Email), zap.Error(err))
		return
	}

	h.logger.Info("Welcome email sent successfully", zap.String("email", payload.Email))
}
