package main

import "os"

func cmdExit(cfg *config) error {
	os.Exit(0)
	return nil
}
