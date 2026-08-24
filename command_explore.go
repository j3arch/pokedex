package main

import (
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	return nil
}
