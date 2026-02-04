package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/harunkilic/rss/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	url := cmd.Args[1]

	dbFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feedfollow: %w", err)
	}

	fmt.Printf("%+v\n", dbFeed)
	return nil
}
