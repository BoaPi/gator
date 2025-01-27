package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username provided")
	}
	fmt.Println(cmd.name, cmd.args)

	if len(cmd.args) > 1 {
		return errors.New("to many arguments")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("New user has been set.")

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.list[cmd.name]; ok {
		err := handler(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("command not found")
}
