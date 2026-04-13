package main

import (
	// "bufio"
	// "fmt"
	// "os"
	"strings"
)

func startRepl() {
	//
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
