package main

import "errors"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must proivde a pokemon name")
	}

}
