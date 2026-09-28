package main

import (
	"context"
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type CommonFeedFollowRow = struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	UserName  string
	FeedName  string
}

func PrintFeedFollowRow[T ~CommonFeedFollowRow](v T) {
	row := CommonFeedFollowRow(v)
	fmt.Println(" - ", row.ID)
	fmt.Println(" - ", row.ID)
	fmt.Println(" - ", row.CreatedAt)
	fmt.Println(" - ", row.UpdatedAt)
	fmt.Println(" - ", row.UserID)
	fmt.Println(" - ", row.FeedID)
	fmt.Println(" - ", row.UserName)
	fmt.Println(" - ", row.FeedName)
	fmt.Println("===========================================================")
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument reguired, received %d", len(cmd.args))
	}

	ctx := context.Background()
	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error: unable to get feed with provided URL(%s): %w", cmd.args[0], err)
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: 1 argument required, received %d", len(cmd.args))
	}

	ctx := context.Background()
	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error: unable to get feed with provided URL(%s): %w", cmd.args[0], err)
	}

	deleteFeedFollowArg := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if _, err := s.db.DeleteFeedFollow(ctx, deleteFeedFollowArg); err != nil {
		return fmt.Errorf("Error: unable to delete the feed follow record: %w", err)
	}

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	followedFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.Name)

	for _, followedFeed := range followedFeeds {
		PrintFeedFollowRow(followedFeed)
	}
	return err
}
