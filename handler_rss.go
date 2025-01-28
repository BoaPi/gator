package main

import (
	"context"
	"fmt"
	"html"

	"github.com/BoaPi/gator/internal/rss"
)

func handlerFetchFeed(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couln't fetch RSS from %s: %w", cmd.Args[0], err)
	}

	// feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	// feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	fmt.Println(html.UnescapeString(feed.Channel.Item[0].Description))
	fmt.Println(feed)
	return nil
}
