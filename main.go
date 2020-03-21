package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		if text == "show" {
			RenderTodos(repo)
		}
		
		args, command := parseInput(text)

		todo := Todo {
			text: args, 
			id: i,
		}
		
		if command == "add" {
			repo, i = AppendTodo(repo, todo)
			fmt.Println("Success added todo with id: ", todo.id)
		}
		
		
		fmt.Print("Tell me what to do: ")
	}
	fmt.Println(repo)
}

func parseInput (text string) (args string, command string) {
	inputArray := strings.Split(text, " ")
	command = inputArray[0]
	args = strings.Join(inputArray[1:], " ")
	return
}