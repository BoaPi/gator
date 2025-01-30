package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BoaPi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	now := time.Now().UTC()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("Following successfully:")
	fmt.Printf(" * Feed:            %s\n", feedFollow.FeedName)
	fmt.Printf(" * User followed:   %s\n", feedFollow.UserName)
	fmt.Println("=====================================")

	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	fmt.Println("Current following:")
	for _, follow := range follows {
		fmt.Printf(" * %s\n", follow.FeedName)
	}

	return nil
}
