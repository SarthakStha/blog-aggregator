package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SarthakStha/gator/internal/database"
	"github.com/SarthakStha/gator/internal/feed"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

func displayPost(post database.Post) {
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("URL: %s\n", post.Url)
	fmt.Printf("Created At: %s\n", post.CreatedAt)
	fmt.Printf("Updated At: %s\n", post.UpdatedAt)
	fmt.Printf("Published At: %s\n", post.PublishedAt.Time)
	fmt.Printf("Description: %s\n", post.Description.String)
	fmt.Println("==================================================================================")
}

func createPosts(s *state, dbFeed database.Feed, item feed.RSSItem) {
	parsedTimeArg := sql.NullTime{}
	if parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate); err != nil {
		fmt.Printf("Unable to parse string in format RFC1123Z, storing as NULL: %s\n", err)
	} else {
		parsedTimeArg.Time = parsedTime
		parsedTimeArg.Valid = true
	}

	ctx := context.Background()
	postArg := database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     item.Title,
		Url:       item.Link,
		Description: sql.NullString{
			String: item.Description,
			Valid:  true,
		},
		PublishedAt: parsedTimeArg,
		FeedID:      dbFeed.ID,
	}

	_, err := s.db.CreatePost(ctx, postArg)
	if err != nil && !strings.Contains(err.Error(), "one_post_per_url") {
		fmt.Printf("Error: unable to create a post: %s\n", err)
	}
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var displayLimit int32 = 2
	if len(cmd.args) == 1 {
		if specifiedLimit, err := strconv.ParseInt(cmd.args[0], 10, 32); err != nil {
			return fmt.Errorf("Error: command expects argument of type int: %w", err)
		} else {
			displayLimit = int32(specifiedLimit)
		}
	}

	ctx := context.Background()
	userPostsArg := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  displayLimit,
	}

	userPosts, err := s.db.GetPostsForUser(ctx, userPostsArg)
	if err != nil {
		fmt.Printf("Error: Unable to get posts from the user: %w", err)
	}

	fmt.Printf("Found %d posts for user: %s\n", len(userPosts), user.Name)
	fmt.Println("==================================================================================")
	for _, post := range userPosts {
		displayPost(post)
	}
	return nil
}
