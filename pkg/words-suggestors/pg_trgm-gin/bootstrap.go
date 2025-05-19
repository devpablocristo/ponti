package pkgsuggester

// Bootstrap receives parameters, configures, creates and starts the suggestion engine.
func Bootstrap(dsn, table, column string) (Suggester, error) {
	// 1. Build configuration from args and env
	opts := []Option{}
	if dsn != "" {
		opts = append(opts, WithDSN(dsn))
	}
	if table != "" {
		opts = append(opts, WithTable(table))
	}
	if column != "" {
		opts = append(opts, WithColumn(column))
	}

	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}

	// 2. Create and start engine
	return newEngine(cfg)
}
