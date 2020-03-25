package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todoRepo TodoRepo
var scanner *bufio.Scanner

func init() {
	todoRepo = TodoRepo{}
	scanner = bufio.NewScanner(os.Stdin)
}

func main() {
	fmt.Print("Tell me what to do: ")

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "quit" || text == "exit" {
			fmt.Println("Bye bye")
			break
		}
		query := parseInput(text)
		todoRepo = executeQuery(todoRepo, query)
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(todoRepo)
}
