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

	programState := &state{
		cfg: &cfg,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	command := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.run(programState, command)
	if err != nil {
		log.Fatal(err)
	}
}
