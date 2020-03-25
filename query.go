package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type QueryArgs struct {
	todoID int
	text   string
}

type Command int

type Query struct {
	command Command
	args    QueryArgs
}

var validCommands = map[string]Command{
	"show":   1,
	"view":   1,
	"add":    2,
	"new":    2,
	"create": 2,
	"edit":   3,
	"update": 3,
	"toggle": 4,
	"done":   4,
	"remove": 5,
	"delete": 5,
	"rm":     5,
}

const (
	show Command = iota + 1
	add
	edit
	toggle
	remove
)

func parseInput(query string) (Query, error) {
	valid, _ := validQuery(query)
	if !valid {
		return Query{command: -1}, errors.New(fmt.Sprint("not a valid query"))
	}

	tokens := strings.Split(query, " ")

	command := getCommand(tokens[0])

	if contains([]int{3, 4, 5}, int(command)) {
		id, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			return Query{command: -1}, errors.New(fmt.Sprint("where is todo id?"))
		}
		return Query{
			command: command,
			args: QueryArgs{
				todoID: int(id),
				text:   strings.Join(tokens[2:], " "),
			},
		}, nil
	} else if command == 2 {
		args := QueryArgs{text: strings.Join(tokens[1:], " ")}
		return Query{command, args}, nil
	}
	return Query{command: command}, nil
}

func getCommand(command string) Command {
	command = strings.ToLower(command)
	if com, ok := validCommands[command]; ok {
		return com
	}
	return -1
}

func validQuery(q string) (bool, error) {
	queryRegex := `(?m)(show|view|add|new|create|edit|update|toggle|done|remove|delete|rm)\s?(\d*)( .*)?`
	return regexp.MatchString(queryRegex, q)
}

func contains(arr []int, n int) bool {
	for _, val := range arr {
		if val == n {
			return true
		}
	}
	return false
}

// add "x" query
// edit ID "x" query
// rm ID query
// toggle ID query
// show ID query
// query = X     Y    Z

// x = command
// y = ID
// z = text

// if command = VALID:
// 	if command  = EDIT || RM || TOGGLE || SHOW:
// 		if ID = existing:
// 		...
// 		else:
// 			goNEXT
// 	if command = Add:
// 		append()
// else:
// 	goNEXT
