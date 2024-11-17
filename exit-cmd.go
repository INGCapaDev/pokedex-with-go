package main

import "os"

func cmdExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
