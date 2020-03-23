package main

import "strings"

type Query struct {
	command string
	args    string
}

func parseInput(query string) Query {
	tokens := strings.Split(query, " ")

	command := tokens[0]
	args := strings.Join(tokens[1:], " ")
	return Query{
		command,
		args,
	}
}
