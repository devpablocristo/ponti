package pkgsuggester

import (
	"fmt"
	"os"
	"regexp"
)

const (
	defaultLimit     = 10
	defaultThreshold = 0.3
)

var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

type Option func(*config)

type config struct {
	DSN       string
	Table     string
	Column    string
	Limit     int
	Threshold float64
	logger    Logger
}

// WithDSN overrides the default database DSN.
func WithDSN(dsn string) Option { return func(c *config) { c.DSN = dsn } }

// WithTable overrides the default table name.
func WithTable(table string) Option { return func(c *config) { c.Table = table } }

// WithColumn overrides the default column name.
func WithColumn(col string) Option { return func(c *config) { c.Column = col } }

// WithLimit overrides the maximum number of suggestions.
func WithLimit(n int) Option { return func(c *config) { c.Limit = n } }

// WithThreshold overrides the pg_trgm similarity threshold.
func WithThreshold(t float64) Option { return func(c *config) { c.Threshold = t } }

// WithLogger sets a custom logger implementation.
func WithLogger(l Logger) Option { return func(c *config) { c.logger = l } }

// newConfig builds the configuration from options and environment, then validates it.
func newConfig(opts ...Option) (*config, error) {
	c := &config{
		DSN:       os.Getenv("DATABASE_DSN"),
		Table:     os.Getenv("SUGGESTER_TABLE"),
		Column:    os.Getenv("SUGGESTER_COLUMN"),
		Limit:     defaultLimit,
		Threshold: defaultThreshold,
		logger:    noopLogger{},
	}
	for _, o := range opts {
		o(c)
	}
	if err := c.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return c, nil
}

// validate ensures all config fields are set and identifiers are safe.
func (c *config) validate() error {
	if c.DSN == "" {
		return fmt.Errorf("DATABASE_DSN is required")
	}
	if c.Limit <= 0 {
		return fmt.Errorf("Limit must be > 0, got %d", c.Limit)
	}
	if c.Threshold < 0 || c.Threshold > 1 {
		return fmt.Errorf("Threshold must be between 0 and 1, got %f", c.Threshold)
	}
	if !validIdentifier.MatchString(c.Table) {
		return fmt.Errorf("invalid table name: %s", c.Table)
	}
	if !validIdentifier.MatchString(c.Column) {
		return fmt.Errorf("invalid column name: %s", c.Column)
	}
	return nil
}
