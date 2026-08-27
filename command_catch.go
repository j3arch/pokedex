package main

import (
	"errors"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name.")
	}

}
