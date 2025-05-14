package wire

import (
	config "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/config"
)

func ProvideConfigLoader() (config.Loader, error) {
	return config.NewConfigLoader()
}
