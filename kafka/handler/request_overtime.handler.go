package handler

import (
	"fmt"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/kafka/payload"
	"github.com/ngdangkietswe/swe-notification-service/mail"
	mailpayload "github.com/ngdangkietswe/swe-notification-service/mail/model"
	"go.uber.org/zap"
)

type RequestOvertimeHandler struct {
	logger *logger.Logger
}

func NewRequestOvertimeHandler(logger *logger.Logger) *RequestOvertimeHandler {
	return &RequestOvertimeHandler{
		logger: logger,
	}
}

// Handle sends an email to the user to notify them of the overtime request.
func (h *RequestOvertimeHandler) Handle(payload *payload.RequestOvertimePayload) {
	emailPayload := mailpayload.MailPayload{
		To:      payload.ApproverEmail,
		Subject: "Overtime Request",
		Body: fmt.Sprintf(
			"Hello %s,\n\nUser <strong>%s<strong> has requested for overtime on <strong>%s</strong> with the following details:\n\n - Total hours: %.1f\n - Reason: %s\n\nPlease review the request and take necessary actions.\n\nBest regards,\nSWE Notification Service Team",
			payload.ApproverEmail,
			payload.Requester,
			payload.Date,
			payload.TotalHours,
			payload.Reason,
		),
	}

	h.logger.Info("Sending email to approver for overtime request", zap.String("email", payload.ApproverEmail))

	if err := mail.SendMail(emailPayload); err != nil {
		h.logger.Error("Failed to send email to approver for overtime request", zap.String("email", payload.ApproverEmail), zap.Error(err))
	}
}
