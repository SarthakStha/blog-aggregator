package main

import (
	"context"
	"github.com/SarthakStha/gator/internal/database"
	"log"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			log.Fatal("Error: current user not found, login or register first")
		}

		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			log.Fatal("Error: current user not found, login or register first")
		}

		return handler(s, cmd, user)
	}
}
