package main

import (
	"context"
	"fmt"

	"github.com/harunkilic/rss/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)

	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}
