package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/BoaPi/gator/internal/config"
	"github.com/BoaPi/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg     *config.Config
	db      *sql.DB
	queries *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("error opening db connection")
		return
	}

	dbQueries := database.New(db)

	programState := &state{
		cfg:     &cfg,
		db:      db,
		queries: dbQueries,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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
