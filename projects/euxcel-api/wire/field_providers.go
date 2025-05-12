package wire

import (
	"errors"

	gorm "github.com/alphacodinggroup/euxcel-backend/pkg/databases/sql/gorm"
	mdw "github.com/alphacodinggroup/euxcel-backend/pkg/http/middlewares/gin"
	ginsrv "github.com/alphacodinggroup/euxcel-backend/pkg/http/servers/gin"

	field "github.com/alphacodinggroup/euxcel-backend/projects/euxcel-api/internal/field"
	lot "github.com/alphacodinggroup/euxcel-backend/projects/euxcel-api/internal/lot"

	gorm "github.com/alphacodinggroup/euxcel-backend/pkg/databases/sql/gorm"
	mdw "github.com/alphacodinggroup/euxcel-backend/pkg/http/middlewares/gin"
	ginsrv "github.com/alphacodinggroup/euxcel-backend/pkg/http/servers/gin"

	field "github.com/alphacodinggroup/euxcel-backend/projects/euxcel-api/internal/field"
)

func ProvideFieldRepository(repo gorm.Repository) (field.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return field.NewRepository(repo), nil
}

// ProvideLotRepository creates a Lot repository instance.
func ProvideLotRepository(repo gorm.Repository) (lot.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return lot.NewRepository(repo), nil
}

// ProvideLotUseCases creates Lot use cases.
func ProvideLotUseCases(repo lot.Repository) lot.UseCases {
	return lot.NewUseCases(repo)
}

// ProvideFieldUseCases wires the Field use cases with repository and Lot service.
func ProvideFieldUseCases(
	repo field.Repository,
	lotUC lot.UseCases,
) field.UseCases {
	return field.NewUseCases(repo, lotUC)
}
(repo field.Repository) field.UseCases {
	return field.NewUseCases(repo)
}

// ProvideFieldHandler creates the HTTP handler for Field endpoints.
func ProvideFieldHandler(
	server ginsrv.Server,
	fieldUC field.UseCases,
	middlewares *mdw.Middlewares,
) *field.Handler {
	return field.NewHandler(server, fieldUC, middlewares)
}
