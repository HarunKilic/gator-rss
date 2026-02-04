package rss

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/harunkilic/rss/internal/database"
	"github.com/lib/pq"
)

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {
	nextFeed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get next feed\n")
	}

	db.MarkFeedFetched(ctx, nextFeed.ID)

	feed, err := FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Failed to fetch feed %v\n", err)
	}

	for _, item := range feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			UpdatedAt:   time.Now().UTC(),
			FeedID:      nextFeed.ID,
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			PublishedAt: publishedAt,
		})

		if err != nil {
			if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
				continue
			}

			log.Printf("Failed to create post: %v\n", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(feed.Channel.Item))

	return nil
}
