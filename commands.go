package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, ok := c.registry[cmd.name]
	if !ok {
		return fmt.Errorf("Error: '%s' command not in registry", cmd.name)
	}
	return cmdHandler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}

func initializeRegistry() *commands {
	registry := &commands{
		make(map[string]func(*state, command) error),
	}
	registry.register("login", handlerLogin)
	registry.register("register", handlerRegister)
	registry.register("reset", handlerReset)
	registry.register("users", handlerListUsers)
	registry.register("agg", handlerAggregation)
	registry.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	registry.register("removefeed", middlewareLoggedIn(handlerRemoveFeed))
	registry.register("feeds", handlerListAllFeeds)
	registry.register("follow", middlewareLoggedIn(handlerFollow))
	registry.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	registry.register("following", middlewareLoggedIn(handlerFollowing))
	registry.register("browse", middlewareLoggedIn(handlerBrowse))

	return registry
}
