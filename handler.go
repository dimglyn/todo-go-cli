package main

import (
	"fmt"
	"strconv"
)

func executeQuery(repo TodoRepo, query Query) TodoRepo {
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
