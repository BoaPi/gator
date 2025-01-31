package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BoaPi/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <number-of-posts> (optional)", cmd.Name)
	}

	var limit string = "2"

	if len(cmd.Args) == 1 {
		limit = cmd.Args[0]
	}

	l, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return fmt.Errorf("couldn't parse limit: %w", err)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(l),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	fmt.Printf("Posts for User: %s\n", user.Name)
	for _, post := range posts {
		fmt.Printf(" * %s\n", post.Title)
	}
	fmt.Println("======================================")

	return nil
}
