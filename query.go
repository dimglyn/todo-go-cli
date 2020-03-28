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

const (
	show Command = iota + 1
	add
	edit
	toggle
	remove
	unknown = -1
)

func parseInput(query string) (Query, error) {
	valid, _ := validQuery(query)
	if !valid {
		return Query{command: unknown}, errors.New(fmt.Sprint("not a valid query"))
	}

	tokens := strings.Split(query, " ")

	command := getCommand(tokens[0])

	if contains([]Command{edit, toggle, remove}, command) {
		id, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			return Query{command: unknown}, errors.New(fmt.Sprint("where is todo id?"))
		}
		return Query{
			command: command,
			args: QueryArgs{
				todoID: int(id),
				text:   strings.Join(tokens[2:], " "),
			},
		}, nil
	} else if command == add {
		args := QueryArgs{text: strings.Join(tokens[1:], " ")}
		return Query{command, args}, nil
	}
	return Query{command: command}, nil
}

func getCommand(command string) Command {
	validCommands := map[string]Command{
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

	command = strings.ToLower(command)
	if com, ok := validCommands[command]; ok {
		return com
	}
	return unknown
}

func validQuery(q string) (bool, error) {
	queryRegex := `((?m)(show|view)$|(?m)(add|new|create)\s(.*)$|(?m)(edit|update)\s(\d*)\s(.*)$|(?m)(toggle|done|remove|delete|rm)\s(\d*)$)`
	return regexp.MatchString(queryRegex, q)
}

func contains(arr []Command, n Command) bool {
	for _, val := range arr {
		if val == n {
			return true
		}
	}
	return false
}
