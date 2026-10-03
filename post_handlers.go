package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/database"
	"github.com/SarthakStha/blog-aggregator/internal/feed"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

func createPosts(s *state, dbFeed database.Feed, item feed.RSSItem) {
	parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
	parsedTimeArg := sql.NullTime{
		Time:  parsedTime,
		Valid: true,
	}

	if err != nil {
		fmt.Printf("Unable to parse string in format RFC1123Z, storing as NULL: %s\n", err)
		parsedTimeArg.Valid = false
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

	if _, err := s.db.CreatePost(ctx, postArg); err != nil {
		if strings.Contains(err.Error(), "posts_url_key") {
			return
		}
		fmt.Printf("Error: Unable to create a post: %s\n", err)
	}
}

func displayPosts(post database.Post) {
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("URL: %s\n", post.Url)
	fmt.Printf("Created At: %s\n", post.CreatedAt)
	fmt.Printf("Updated At: %s\n", post.UpdatedAt)
	fmt.Printf("Published At: %s\n", post.PublishedAt.Time)
	fmt.Printf("Description: %s....\n", post.Description.String[:200])
	fmt.Println("==================================================================================")
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var displayLimit int32 = 2
	if len(cmd.args) != 0 {
		i, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error: Command expects argument of type int: %w", err)
		}
		displayLimit = int32(i)
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

	for _, post := range userPosts {
		displayPosts(post)
	}
	return nil
}
