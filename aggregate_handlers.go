package main

import (
	"context"
	"fmt"
	"github.com/SarthakStha/gator/internal/feed"
	"time"
)

func handlerAggregation(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument required, received %d", len(cmd.args))
	}

	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error: unable to parse duration string: %w", err)
	}
	fmt.Println("Collecting feeds every", interval)

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Println(err)
			break
		}
	}
	return nil
}

// Here we are dealing with two types of feed
// dbFeed is the the entry stored in the database
// it contains the feed name at the title
// content feed is the data extracted from the URL
func scrapeFeeds(s *state) error {
	ctx := context.Background()
	dbFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error: unable to get the next feed: %w", err)
	}

	dbFeed, err = s.db.MarkFeedFetched(ctx, dbFeed.ID)
	if err != nil {
		return fmt.Errorf("Error: unable to update the feed: %w", err)
	}

	contentFeeds, err := feed.FetchFeed(ctx, dbFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println("Posting feed from::", dbFeed.Name)
	for _, contentFeed := range contentFeeds.Channel.Item {
		createPosts(s, dbFeed, contentFeed)
	}
	return nil
}
