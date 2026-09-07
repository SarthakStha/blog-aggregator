package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument required, received %d", len(cmd.args))
	}

	if err := s.currConfig.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("user: %s has been set.\n", cmd.args[0])
	return nil
}
