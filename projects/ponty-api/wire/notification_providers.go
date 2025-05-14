package wire

import (
	"errors"

	mdw "github.com/alphacodinggroup/euxcel-backend/pkg/http/middlewares/gin"
	gin "github.com/alphacodinggroup/euxcel-backend/pkg/http/servers/gin"
	smtp "github.com/alphacodinggroup/euxcel-backend/pkg/notification/smtp"

	notification "github.com/alphacodinggroup/euxcel-backend/projects/euxcel-api/internal/notification"
)

func ProvideNotificationSmtpService(smtp smtp.Service) (notification.SmtpService, error) {
	if smtp == nil {
		return nil, errors.New("smtp service cannot be nil")
	}
	return notification.NewSmtpService(smtp), nil
}

func ProvideNotificationUseCases(
	ssrv notification.SmtpService,
) notification.UseCases {
	return notification.NewUseCases(ssrv)
}

func ProvideNotificationHandler(
	server gin.Server,
	usecases notification.UseCases,
	middlewares *mdw.Middlewares,
) *notification.Handler {
	return notification.NewHandler(server, usecases, middlewares)
}
