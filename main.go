package main

import (
	"fmt"
	"log"

	"github.com/BoaPi/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("BoaPi")
	if err != nil {
		log.Fatal("user name could not be set")
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	fmt.Printf("Read config agian: %v\n", cfg)
}
