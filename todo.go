package main

import (
	"errors"
	"fmt"
)

type Todo struct {
	text string
	id   int
	done bool
}

type TodoRepo []Todo

func (tr TodoRepo) String() (todos string) {
	for _, todo := range tr {
		done := ""
		if todo.done {
			done = "\u2713"
		}
		todos += fmt.Sprintf("%v - %v %v\n", todo.id, todo.text, done)
	}
	return todos
}

func AppendTodo(tr TodoRepo, todo Todo) (TodoRepo, int) {
	tr = append(tr, todo)
	return tr, todo.id + 1
}

func EditTodoText(tr TodoRepo, todoID int, updatedText string) (error, TodoRepo, Todo) {
	err, index := findIndex(tr, todoID)
	if err != nil {
		return err, tr, Todo{}
	}

	tr[index].text = updatedText

	return nil, tr, tr[index]
}

func ToggleDone(tr TodoRepo, todoID int) (TodoRepo, Todo, error) {
	err, index := findIndex(tr, todoID)
	if err != nil {
		return tr, Todo{}, err
	}

	tr[index].done = !tr[index].done

	return tr, tr[index], err
}

func RemoveTodo(tr TodoRepo, todoID int) (error, TodoRepo) {
	err, selectedIndexID := findIndex(tr, todoID)
	if err != nil {
		return err, tr
	}

	tr = append(tr[:selectedIndexID], tr[selectedIndexID+1:]...)

	return nil, tr
}

func findIndex(tr TodoRepo, todoID int) (error, int) {
	for i, t := range tr {
		if t.id == todoID {
			return nil, i
		}
	}

	return errors.New(fmt.Sprint("Did not find todo with id: ", todoID)), -1
}
