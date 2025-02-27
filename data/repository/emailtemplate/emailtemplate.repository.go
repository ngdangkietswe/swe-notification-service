package emailtemplate

import (
	"context"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/emailtemplate"
)

type emailTemplateRepository struct {
	entClient *ent.Client
}

func (e emailTemplateRepository) FindByKey(ctx context.Context, key string) (*ent.EmailTemplate, error) {
	return e.entClient.EmailTemplate.Query().Where(emailtemplate.TemplateKey(key)).First(ctx)
}

func NewEmailTemplateRepository(entClient *ent.Client) IEmailTemplateRepository {
	return &emailTemplateRepository{
		entClient: entClient,
	}
}
