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
		query := parseInput(text)
		query.command = strings.ToLower(query.command)
		todoRepo = execute(todoRepo, query)
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(todoRepo)
}

func execute(repo TodoRepo, query Query) TodoRepo {

	switch query.command {
	case "show", "view":
		fmt.Println(repo)
		break
	case "add", "new", "create":
		repo, _ = newTodo(repo, query.args)
		break
	case "edit", "update":
		query := parseInput(query.args)
		if id, err := strconv.ParseInt(query.command, 10, 32); err == nil {
			repo, _ = updateTodo(repo, int(id), query.args)
			break
		}
	case "toggle", "done":
		if index, err := strconv.ParseInt(query.args, 10, 32); err == nil {
			repo = toggleDone(repo, int(index))
			break
		}
	case "remove", "delete", "rm":
		if removeIndex, err := strconv.ParseInt(query.args, 10, 32); err == nil {
			repo = deleteTodo(repo, int(removeIndex))
			break
		}
	default:
		fmt.Println("Sorry cant do that")
	}

	return repo
}
