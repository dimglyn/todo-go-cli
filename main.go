package main

import (
	"bufio"
	"flag"
	"fmt"
)

var (
	Token    string
	todoRepo TodoRepo
	scanner  *bufio.Scanner
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	todoRepo = TodoRepo{}
}

func main() {

	fmt.Println(Token)

	// for scanner.Scan() {

	// 	text := strings.TrimSpace(scanner.Text())
	// 	if text == "quit" || text == "exit" {
	// 		fmt.Println("Bye bye")
	// 		break
	// 	}
	// 	query, err := parseInput(text)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		fmt.Print("Tell me what to do: ")
	// 		continue
	// 	}
	// 	todoRepo = handle(todoRepo, query)
	// 	fmt.Print("Tell me what to do: ")
	// }
	// fmt.Println(todoRepo)
}
