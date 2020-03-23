package main

import (
	"fmt"
)

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
