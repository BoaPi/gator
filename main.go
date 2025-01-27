package main

import (
	"log"
	"os"

	"github.com/BoaPi/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	s := state{
		cfg: &cfg,
	}

	commands := commands{
		list: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("no command provided")
	}

	command := command{
		name: args[1],
		args: args[2:],
	}

	err = commands.run(&s, command)
	if err != nil {
		log.Fatal(err)
	}
}
