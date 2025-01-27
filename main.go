package main

import (
	"fmt"
	"log"

	"github.com/BoaPi/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error while reading config file")
	}

	cfg.SetUser("BoaPi")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("error while reading config file")
	}

	fmt.Println(cfg.DBUrl)
	fmt.Println(cfg.CurrentUserName)
}
