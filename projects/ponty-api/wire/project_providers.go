package wire

import (
	"errors"

	gorm "github.com/alphacodinggroup/ponti-backend/pkg/databases/sql/gorm"
	mdw "github.com/alphacodinggroup/ponti-backend/pkg/http/middlewares/gin"
	ginsrv "github.com/alphacodinggroup/ponti-backend/pkg/http/servers/gin"

	project "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/project"
)

func ProvideProjectRepository(repo gorm.Repository) (project.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return project.NewRepository(repo), nil
}

func ProvideProjectUseCases(repo project.Repository) project.UseCases {
	return project.NewUseCases(repo)
}

func ProvideProjectHandler(server ginsrv.Server, usecases project.UseCases, middlewares *mdw.Middlewares) *project.Handler {
	return project.NewHandler(server, usecases, middlewares)
}
