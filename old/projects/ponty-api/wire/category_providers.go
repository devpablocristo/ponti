package wire

import (
	"errors"

	gorm "github.com/alphacodinggroup/ponti-backend/pkg/databases/sql/gorm"
	mdw "github.com/alphacodinggroup/ponti-backend/pkg/http/middlewares/gin"
	ginsrv "github.com/alphacodinggroup/ponti-backend/pkg/http/servers/gin"

	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/category"
)

func ProvideCategoryRepository(repo gorm.Repository) (category.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return category.NewRepository(repo), nil
}

func ProvideCategoryUseCases(
	repo category.Repository,
) category.UseCases {
	return category.NewUseCases(repo)
}

func ProvideCategoryHandler(
	server ginsrv.Server,
	usecases category.UseCases,
	middlewares *mdw.Middlewares,
) *category.Handler {
	return category.NewHandler(server, usecases, middlewares)
}
