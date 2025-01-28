package main

import (
	"context"
	"fmt"

	"github.com/BoaPi/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couln't fetch RSS from %s: %w", cmd.Args[0], err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
