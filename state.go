package main

import (
	"database/sql"
	"github.com/SarthakStha/blog-aggregator/internal/config"
	"github.com/SarthakStha/blog-aggregator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func initializeState() (*state, error) {
	userConfig, err := config.Read()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", userConfig.DbURL)
	if err != nil {
		return nil, err
	}
	dbQueries := database.New(db)

	// Here commana is required because go
	// automatically adds a semicolon at the end
	// of the expression.
	return &state{
		&userConfig,
		dbQueries,
	}, nil
}
