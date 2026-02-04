package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		log.Fatalf("couldn't reset users: %v", err)
	}

	fmt.Println("Users database reset successfully!")
	return nil
}
