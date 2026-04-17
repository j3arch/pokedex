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

func getCommands() map[string]cliCommand {
	return nil
}
