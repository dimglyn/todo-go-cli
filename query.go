package main

import (
	"regexp"
	"strings"
)

type QueryArgs struct {
	todoID int
	text   string
}

type Command int

type Query struct {
	command Command
	args    string
}

const validCommands = [string]Command{
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

func parseInput(query string) Query {

	tokens := strings.Split(query, " ")

	command := getCommand(tokens[0])
	args := strings.Join(tokens[1:], " ")
	return Query{
		command,
		args,
	}
}

func getCommand(command string) Command {
	command = strings.ToLower(command)
	if com, ok := validCommands[command]; ok {
		return com
	} else {
		return -1
	}
}

func validQuery(q string) (bool, error) {
	queryRegex := `(?m)(show|view|add|new|create|edit|update|toggle|done|remove|delete|rm)\s+(\d*)( .*)?`
	return regexp.MatchString(queryRegex, q)
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
