package pkgsuggester

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// engine implements Suggester using GORM with pg_trgm extension.
type engine struct {
	db     *gorm.DB
	table  string
	column string
	limit  int
	logger Logger
}

// newEngine creates an engine from a validated config.
func newEngine(cfg *config) (Suggester, error) {
	// open GORM connection
	gormDB, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect with GORM: %w", err)
	}
	// ensure database is reachable
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("retrieve generic DB: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}
	// set similarity threshold
	if err := gormDB.WithContext(ctx).
		Exec("SET pg_trgm.similarity_threshold = ?", cfg.Threshold).
		Error; err != nil {
		return nil, fmt.Errorf("set threshold: %w", err)
	}
	// return initialized engine
	return &engine{
		db:     gormDB,
		table:  cfg.Table,
		column: cfg.Column,
		limit:  cfg.Limit,
		logger: cfg.logger,
	}, nil
}

// NewSuggester builds config from options and returns a ready Suggester.
func NewSuggester(opts ...Option) (Suggester, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}
	return newEngine(cfg)
}

// Suggest returns up to limit suggestions matching the prefix, ordered by similarity.
func (e *engine) Suggest(ctx context.Context, query string) ([]Suggestion, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, nil
	}
	// build SQL with safe identifiers
	sql := fmt.Sprintf(
		`SELECT id, %s AS text
		 FROM %s
		 WHERE %s ILIKE ?
		 ORDER BY similarity(%s, ?) DESC
		 LIMIT %d`,
		e.column, e.table, e.column, e.column, e.limit,
	)

	var results []Suggestion
	start := time.Now()
	err := e.db.WithContext(ctx).
		Raw(sql, q+"%", q).
		Scan(&results).
		Error
	if err != nil {
		e.logger.Error("suggest query failed", err)
		return nil, fmt.Errorf("suggest query: %w", err)
	}
	e.logger.Debug(fmt.Sprintf("suggest latency: %s", time.Since(start)))
	return results, nil
}

// Close gracefully closes the underlying database connection.
func (e *engine) Close() error {
	sqlDB, err := e.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Health checks the liveliness of the database connection.
func (e *engine) Health(ctx context.Context) error {
	sqlDB, err := e.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
