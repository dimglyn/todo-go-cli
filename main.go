package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
)
		type Todo struct {
	text string
	id int
}

type TodoRepo []Todo

func main() {
	repo := TodoRepo{}
	i := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Tell me what to do: ")
	for scanner.Scan() {
	
		text := scanner.Text()
		
		if text == "quit" {
			break
		}
		if(text == "show") {
			renderTodos(repo)
		}
		
		inputArray := strings.Split(text, " ")
		command := inputArray[0]
		
		args := strings.Join(inputArray[1:], " ")
		
		todo := Todo {
			text: args, 
			id: i,
		}
		
		if command == "add"{
			repo, i = appendTodo(repo, todo)
			fmt.Println("Success added todo with id: ", todo.id)
		}
		
		
		fmt.Print("Tell me what to do: ")
		
	}
	fmt.Println(repo)
}

func appendTodo (tr TodoRepo, todo Todo) (TodoRepo, int) {
	tr = append(tr, todo)
	return tr, todo.id + 1
}

func renderTodos (tr TodoRepo) {
	fmt.Println(tr)
}