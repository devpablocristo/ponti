package pkgsuggester

import "context"

// Suggestion represents one match result.
type Suggestion struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Suggester defines the methods of the suggestion engine.
type Suggester interface {
	// Suggest returns suggestions for the given query.
	Suggest(ctx context.Context, query string) ([]Suggestion, error)
	// Close releases all resources associated with the engine.
	Close() error
	// Health performs a health check of dependencies.
	Health(ctx context.Context) error
}

// Logger provides Debug and Error logging methods.
type Logger interface {
	Debug(msg string)
	Error(msg string, err error)
}

// noopLogger is a no-operation logger used by default.
type noopLogger struct{}

func (noopLogger) Debug(_ string)          {}
func (noopLogger) Error(_ string, _ error) {}
