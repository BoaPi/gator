package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/BoaPi/gator/internal/database"
	"github.com/BoaPi/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Errorf("usage: %s <1s, 1m, 1h>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid time %s: %w", cmd.Args[0], err)
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feed to fetch: %w", err)
	}

	rss, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	fmt.Printf("Feed %s fetched:", feed.Name)
	for _, item := range rss.Channel.Item {
		fmt.Printf(" * %s\n", item.Title)
	}
	fmt.Println("================================")

	return nil
}
