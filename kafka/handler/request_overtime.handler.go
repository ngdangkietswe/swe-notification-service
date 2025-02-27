package handler

import (
	"context"
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/constants"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
	"github.com/ngdangkietswe/swe-notification-service/data/repository/emailtemplate"
	"github.com/ngdangkietswe/swe-notification-service/kafka/payload"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"strings"
)

type RequestOvertimeHandler struct {
	logger            *logger.Logger
	emailTemplateRepo emailtemplate.IEmailTemplateRepository
}

func NewRequestOvertimeHandler(logger *logger.Logger, emailTemplateRepo emailtemplate.IEmailTemplateRepository) *RequestOvertimeHandler {
	return &RequestOvertimeHandler{
		logger:            logger,
		emailTemplateRepo: emailTemplateRepo,
	}
}

// Handle sends an email to the user to notify them of the overtime request.
func (h *RequestOvertimeHandler) Handle(payload *payload.RequestOvertimePayload) {
	emailTemplate, err := h.emailTemplateRepo.FindByKey(context.Background(), constants.EmailTemplateKeyOvertimeRequest)
	if err != nil {
		h.logger.Error(
			lo.Ternary(ent.IsNotFound(err), "Email template not found", "Failed to get Email template"),
			zap.String("error", err.Error()),
		)
		return
	}

	body := strings.NewReplacer(
		"{approver}", payload.Approver,
		"{requester}", payload.Requester,
		"{date}", payload.Date,
		"{overtime_hours}", fmt.Sprintf("%.2f", payload.TotalHours),
		"{reason}", payload.Reason,
	).Replace(emailTemplate.Body)

	emailPayload := mailpayload.MailPayload{
		To:      payload.ApproverEmail,
		Subject: emailTemplate.Subject,
		Body:    body,
	}

	h.logger.Info("Sending email to approver for overtime request", zap.String("email", payload.ApproverEmail))

	if err = mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send email to approver for overtime request", zap.String("email", payload.ApproverEmail), zap.Error(err))
		return
	}

	h.logger.Info("Overtime request email sent successfully", zap.String("recipient", payload.ApproverEmail))
}
