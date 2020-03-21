package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	repo := TodoRepo{}
	i := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Tell me what to do: ")

	for scanner.Scan() {
		text := scanner.Text()
		repo = execute(repo, text, i)
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(repo)
}

func execute(repo TodoRepo, query string, i int) TodoRepo {
	command, args := parseInput(query)

	switch command {
	case "show":
		RenderTodos(repo)
	case "add":
		repo, i = addTodo(repo, i, args)
	case "edit":
		idString, updatedText := parseInput(args)
		if id, err := strconv.ParseInt(idString, 10, 32); err == nil {
			repo, _ = updateTodo(repo, int(id), updatedText)
		}
	case "remove":
		if removeIndex, err := strconv.ParseInt(args, 10, 32); err == nil {
			repo = deleteTodo(repo, int(removeIndex))
		}
	case "quit":
		fmt.Println("Bye bye")
		break
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

func addTodo(repo TodoRepo, i int, args string) (TodoRepo, int) {
	todo := Todo{
		text: args,
		id:   i,
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
	err, repo, updatedTodo := EditTodoByID(repo, id, updatedText)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("updated todo ", updatedTodo)
	}
	return repo, updatedTodo
}
