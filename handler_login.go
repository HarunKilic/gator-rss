package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		log.Fatalf("couldn't find user: %v", name)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User switched successfully! %s\n", user.Name)
	return nil
}
