package main

import (
	"context"
	"errors"

	"github.com/harunkilic/rss/internal/database"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.registeredCommands[cmd.Name]; ok {
		return f(s, cmd)
	}
	return errors.New("command not found")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.New("you must be logged in to use this command")
		}
		return handler(s, cmd, user)
	}
}
