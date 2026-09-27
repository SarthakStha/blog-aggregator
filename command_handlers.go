package main

import (
	"context"
	"fmt"
	"github.com/SarthakStha/blog-aggregator/internal/feed"
)

func handlerAggregation(s *state, cmd command) error {
	ctx := context.Background()
	// test URL
	url := "https://www.wagslane.dev/index.xml"

	feed, err := feed.FetchFeed(ctx, url)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
