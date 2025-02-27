package cdcauthusers

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/kafka"
)

type ICdcAuthUsersRepository interface {
	DeleteById(ctx context.Context, id uuid.UUID) error
	UpsertByCdcEventMsg(ctx context.Context, cdcEventMsg *kafka.CdcEventMsg) error
}
