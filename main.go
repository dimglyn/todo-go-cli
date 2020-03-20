package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	text string
	id int
}

type TodoRepo []Todo

func main() {
	todos := TodoRepo{}
	i := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Tell me what to do: ")
    text := scanner.Text()

		if text == "skase" {
			break
		}

		todo := Todo {
			text: text,
			id: i,
		}

		todos = addTodo(todos, todo)

		i++
	}
	fmt.Println(todos)
}

func addTodo (tr TodoRepo, todo Todo) TodoRepo {
	tr = append(tr, todo)
	return tr
}