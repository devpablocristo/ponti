package wire

import (
	"errors"

	gorm "github.com/alphacodinggroup/ponti-backend/pkg/databases/sql/gorm"
	mdw "github.com/alphacodinggroup/ponti-backend/pkg/http/middlewares/gin"
	ginsrv "github.com/alphacodinggroup/ponti-backend/pkg/http/servers/gin"

	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/authe"
	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/config"
	"github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/item"
)

// ProvideItemRepository inyecta la implementación de Repository para Item.
func ProvideItemRepository(repo gorm.Repository) (item.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return item.NewRepository(repo), nil
}

// ProvideItemUseCases inyecta las dependencias requeridas por la capa de casos de uso de Item.
func ProvideItemUseCases(
	repo item.Repository,
	cfg config.Loader,
	au authe.UseCases,
) item.UseCases {
	return item.NewUseCases(repo, cfg, au)
}

// ProvideItemHandler inyecta las dependencias para crear el Handler de Item.
func ProvideItemHandler(
	server ginsrv.Server,
	usecases item.UseCases,
	middlewares *mdw.Middlewares,
) *item.Handler {
	return item.NewHandler(server, usecases, middlewares)
}
