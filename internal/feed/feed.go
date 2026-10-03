package feed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func (feed *RSSFeed) normalizeText() {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
}

func (feed *RSSFeed) String() string {
	feedText := fmt.Sprintf("Title: %s, Link: %s\nDescription: %s\n", feed.Channel.Title, feed.Channel.Link, feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		feedText += fmt.Sprint("\t", &item)
	}
	return feedText
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (item *RSSItem) normalizeText() {
	item.Title = html.UnescapeString(item.Title)
	item.Description = html.UnescapeString(item.Description)
}

func (item *RSSItem) String() string {
	return fmt.Sprintf("Title: %s, \nLink: %s\nDescription: %s\nPublished Date: %s\n",
		item.Title, item.Link, item.Description, item.PubDate)
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	dataFeed := &RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return dataFeed, fmt.Errorf("Error: unable to create request: %w", err)
	}

	req.Header.Add("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return dataFeed, fmt.Errorf("Error: unable to make a req: %w", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return dataFeed, fmt.Errorf("Error: unable to read the data: %w", err)
	}

	if err := xml.Unmarshal(data, dataFeed); err != nil {
		return dataFeed, fmt.Errorf("Error: unable to parse the data: %w", err)
	}

	dataFeed.normalizeText()
	for i := range len(dataFeed.Channel.Item) {
		(&dataFeed.Channel.Item[i]).normalizeText()
	}
	return dataFeed, nil
}
