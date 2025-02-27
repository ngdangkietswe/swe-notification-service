package cdcauthusers

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/kafka"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/cdcauthusers"
)

type cdcAuthUsersRepository struct {
	entClient *ent.Client
}

// DeleteById deletes a record by its ID.
func (c cdcAuthUsersRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return c.entClient.CdcAuthUsers.DeleteOneID(id).Exec(ctx)
}

// UpsertByCdcEventMsg upserts a record by a CDC event message.
func (c cdcAuthUsersRepository) UpsertByCdcEventMsg(ctx context.Context, cdcEventMsg *kafka.CdcEventMsg) error {
	data := cdcEventMsg.After
	cdcAuthUsersId := data["id"].(string)

	return c.entClient.CdcAuthUsers.Create().
		SetID(uuid.MustParse(cdcAuthUsersId)).
		SetUsername(data["username"].(string)).
		SetEmail(data["email"].(string)).
		OnConflictColumns(cdcauthusers.FieldID).
		UpdateNewValues().
		Exec(ctx)
}

func NewCdcAuthUsersRepository(entClient *ent.Client) ICdcAuthUsersRepository {
	return &cdcAuthUsersRepository{
		entClient: entClient,
	}
}
