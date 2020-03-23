package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		text := scanner.Text()
		if text == "quit" || text == "exit" {
			fmt.Println("Bye bye")
			break
		}
		todoRepo = execute(todoRepo, text)
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(todoRepo)
}

func execute(repo TodoRepo, query string) TodoRepo {
	command, args := parseInput(query)
	command = strings.ToLower(command)

	switch command {
	case "show", "view":
		fmt.Println(repo)
	case "add", "new", "create":
		repo, _ = newTodo(repo, args)
	case "edit", "update":
		idString, updatedText := parseInput(args)
		if id, err := strconv.ParseInt(idString, 10, 32); err == nil {
			repo, _ = updateTodo(repo, int(id), updatedText)
		}
	case "toggle", "done":
		if index, err := strconv.ParseInt(args, 10, 32); err == nil {
			repo = toggleDone(repo, int(index))
		}
	case "remove", "delete", "rm":
		if removeIndex, err := strconv.ParseInt(args, 10, 32); err == nil {
			repo = deleteTodo(repo, int(removeIndex))
		}
	default:
		fmt.Println("Sorry cant do that")
	}

	return repo
}
