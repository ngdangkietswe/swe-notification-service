package repository

import (
	"github.com/ngdangkietswe/swe-notification-service/data/datasource"
	"github.com/ngdangkietswe/swe-notification-service/data/repository/cdcauthusers"
	"github.com/ngdangkietswe/swe-notification-service/data/repository/emailtemplate"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	datasource.NewEntClient,
	cdcauthusers.NewCdcAuthUsersRepository,
	emailtemplate.NewEmailTemplateRepository,
)
