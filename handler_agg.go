package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/BoaPi/gator/internal/database"
	"github.com/BoaPi/gator/internal/rss"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration %s: %w", cmd.Args[0], err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("couldn't get feed to fetch: %w", err)
		return
	}

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("couldn't fetch feed: %w", err)
		return
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("couldn't mark feed %s fetched: %w", feed.Name, err)
		return
	}

	fmt.Printf("Feed %s fetched:\n", feed.Name)

	for _, item := range feedData.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("couldn't parse published date: %w\n", err)
			fmt.Printf("%s\n", item.PubDate)
			continue
		}

		var title sql.NullString
		title.String = item.Title

		if item.Title != "" {
			title.Valid = true
		} else {
			title.Valid = false
		}

		var description sql.NullString
		description.String = item.Description

		if item.Description != "" {
			description.Valid = true
		} else {
			description.Valid = false
		}

		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt.UTC(),
			FeedID:      feed.ID,
		})
	}
	fmt.Println("================================")
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
	fmt.Println("================================")
}
