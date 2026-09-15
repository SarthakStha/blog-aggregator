package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument required, received %d", len(cmd.args))
	}

	ctx := context.Background()
	if _, err := s.db.GetUser(ctx, cmd.args[0]); err != nil {
		return fmt.Errorf("Error while trying to fetch user '%s': %s", cmd.args[0], err)
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("Logged in as user '%s'\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument required, received %d", len(cmd.args))
	}

	ctx := context.Background()
	if _, err := s.db.GetUser(ctx, cmd.args[0]); !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("Error: user '%s' already exists", cmd.args[0])
	}

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	if _, err := s.db.CreateUser(ctx, user); err != nil {
		return err
	}

	fmt.Printf("User '%s' has been created.\n", cmd.args[0])
	return handlerLogin(s, cmd)
}
