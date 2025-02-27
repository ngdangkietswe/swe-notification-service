package handler

import (
	"context"
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

type ReplyOvertimeHandler struct {
	logger            *logger.Logger
	emailTemplateRepo emailtemplate.IEmailTemplateRepository
}

func NewReplyOvertimeHandler(logger *logger.Logger, emailTemplateRepo emailtemplate.IEmailTemplateRepository) *ReplyOvertimeHandler {
	return &ReplyOvertimeHandler{
		logger:            logger,
		emailTemplateRepo: emailTemplateRepo,
	}
}

// Handle sends an email to the user to notify them of the overtime request reply.
func (h *ReplyOvertimeHandler) Handle(payload *payload.ReplyOvertimePayload) {
	emailTemplate, err := h.emailTemplateRepo.FindByKey(context.Background(), constants.EmailTemplateKeyOvertimeReply)
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
		"{status}", lo.Ternary(payload.IsApproved, "Approved", "Rejected"),
		"{reason}", payload.Reason,
	).Replace(emailTemplate.Body)

	emailPayload := mailpayload.MailPayload{
		To:      payload.RequesterEmail,
		Subject: emailTemplate.Subject,
		Body:    body,
	}

	h.logger.Info("Sending email to requester for overtime request reply", zap.String("email", payload.RequesterEmail))

	if err = mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send email to requester for overtime request reply", zap.String("email", payload.RequesterEmail), zap.Error(err))
		return
	}

	h.logger.Info("Overtime reply email sent successfully", zap.String("recipient", payload.RequesterEmail))
}
