package payload

type RequestOvertimePayload struct {
	Date          string  `json:"date"`
	Requester     string  `json:"requester"`
	ApproverEmail string  `json:"approver_email"`
	TotalHours    float32 `json:"total_hours"`
	Reason        string  `json:"reason"`
}
