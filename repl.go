package main

import (
	// "bufio"
	// "fmt"
	// "os"
	"strings"
)

func startRepl() {
	// setup repl
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return nil
}
