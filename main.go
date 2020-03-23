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

func parseInput(text string) (command string, args string) {
	inputArray := strings.Split(text, " ")
	command = inputArray[0]
	args = strings.Join(inputArray[1:], " ")
	return
}

func newTodo(repo TodoRepo, args string) (TodoRepo, int) {
	i := len(repo)
	todo := Todo{
		text: args,
		id:   i,
		done: false,
	}

	fmt.Println("Success added todo with id: ", todo.id)
	return AppendTodo(repo, todo)
}

func deleteTodo(repo TodoRepo, index int) TodoRepo {
	err, repo := RemoveTodo(repo, index)
	if err == nil {
		fmt.Println("Removed todo with id: ", index)
	} else {
		fmt.Println(err)
	}
	return repo
}

func updateTodo(repo TodoRepo, id int, updatedText string) (TodoRepo, Todo) {
	err, repo, updatedTodo := EditTodoText(repo, id, updatedText)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("updated todo ", updatedTodo)
	}
	return repo, updatedTodo
}

func toggleDone(repo TodoRepo, index int) TodoRepo {
	err, repo, _ := ToggleDone(repo, index)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Marked as done!")
	}
	return repo
}
