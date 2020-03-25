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
	fmt.Print("Tell me what to do: ")
}

func main() {
	for scanner.Scan() {

		text := strings.TrimSpace(scanner.Text())
		if text == "quit" || text == "exit" {
			fmt.Println("Bye bye")
			break
		}
		query, err := parseInput(text)
		if err != nil {
			fmt.Println(err)
			fmt.Print("Tell me what to do: ")
			continue
		}
		todoRepo = executeQuery(todoRepo, query)
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(todoRepo)
}
