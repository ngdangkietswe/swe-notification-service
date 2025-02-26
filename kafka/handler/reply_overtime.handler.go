package handler

import (
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/kafka/payload"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ReplyOvertimeHandler struct {
	logger *logger.Logger
}

func NewReplyOvertimeHandler(logger *logger.Logger) *ReplyOvertimeHandler {
	return &ReplyOvertimeHandler{
		logger: logger,
	}
}

// Handle sends an email to the user to notify them of the overtime request reply.
func (h *ReplyOvertimeHandler) Handle(payload *payload.ReplyOvertimePayload) {
	emailPayload := mailpayload.MailPayload{
		To:      payload.RequesterEmail,
		Subject: "Overtime Request Reply",
		Body: fmt.Sprintf(
			"Hello %s,\n\nUser <strong>%s</strong> has replied to your overtime request on <strong>%s</strong> with the following details:\n\n - Status: %s\n - Reason: %s\n\nBest regards,\nSWE Notification Service Team",
			payload.RequesterEmail,
			payload.Approver,
			payload.Date,
			lo.Ternary(payload.IsApproved, "Approved", "Rejected"),
			payload.Reason,
		),
	}

	h.logger.Info("Sending email to requester for overtime request reply", zap.String("email", payload.RequesterEmail))

	if err := mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send email to requester for overtime request reply", zap.String("email", payload.RequesterEmail), zap.Error(err))
	}
}
