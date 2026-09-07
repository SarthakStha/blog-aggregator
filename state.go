package main

import (
	"github.com/SarthakStha/blog-aggregator/internal/config"
)

type state struct {
	currConfig *config.Config
}

func initializeState() (*state, error) {
	userConfig, err := config.Read()
	if err != nil {
		return nil, err
	}
	// Here commana is required because go
	// automatically adds a semicolon at the end
	// of the expression.
	return &state{
		&userConfig,
	}, nil
}
