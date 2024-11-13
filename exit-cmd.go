package main

import "os"

func cmdExit() error {
	os.Exit(0)
	return nil
}
