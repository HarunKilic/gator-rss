package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/harunkilic/rss/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	dbFeed, err := s.db.GetFeedByUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
	})
	fmt.Printf("Feed: %s .... User: %s\n", dbFeed.Name, user.Name)
	return nil
}
