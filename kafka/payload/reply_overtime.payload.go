package payload

type ReplyOvertimePayload struct {
	Date           string `json:"date"`
	Approver       string `json:"approver"`
	Requester      string `json:"requester"`
	RequesterEmail string `json:"requester_email"`
	IsApproved     bool   `json:"is_approved"`
	Reason         string `json:"reason"`
}
