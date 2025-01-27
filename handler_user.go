package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BoaPi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.queries.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User %s not registered", name)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	now := time.Now()

	user, err := s.queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User %s created.", user.Name)
	fmt.Printf("%v", user)

	return nil
}
