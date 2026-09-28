package main

import (
	"context"
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func printFeed(feed database.ListAllFeedsRow) {
	fmt.Println(" - ID: ", feed.ID)
	fmt.Println(" - Created At: ", feed.CreatedAt)
	fmt.Println(" - Updated At: ", feed.UpdatedAt)
	fmt.Println(" - Name: ", feed.Name)
	fmt.Println(" - URL: ", feed.Url)
	fmt.Println(" - User ID: ", feed.UserID)
	fmt.Println(" - Username: ", feed.UserName)
	fmt.Println("========================================================")
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Error: 2 arguments required, received %d", len(cmd.args))
	}

	feedArg := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	ctx := context.Background()
	feed, err := s.db.CreateFeed(ctx, feedArg)
	if err != nil {
		return fmt.Errorf("Error: unable to create feed: %w", err)
	}

	feedFollowsArg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowRow, err := s.db.CreateFeedFollow(ctx, feedFollowsArg)
	if err != nil {
		return fmt.Errorf("Error: unable to insert into feed_follow: %w", err)
	}

	PrintFeedFollowRow(feedFollowRow)
	return nil
}

func handlerListAllFeeds(s *state, cmd command) error {
	ctx := context.Background()
	allFeeds, err := s.db.ListAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Error: unable to query feeds: %w", err)
	}

	for _, feed := range allFeeds {
		printFeed(feed)
	}
	return nil
}
