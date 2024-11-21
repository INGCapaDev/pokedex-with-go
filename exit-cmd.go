package main

import (
	"os"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func cmdExit(cfg *config.TConfig, args ...string) error {
	os.Exit(0)
	return nil
}
