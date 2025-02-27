package emailtemplate

import (
	"context"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
)

type IEmailTemplateRepository interface {
	FindByKey(ctx context.Context, key string) (*ent.EmailTemplate, error)
}
