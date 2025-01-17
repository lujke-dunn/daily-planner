package commands 

import (
	"src/list"
	"fmt"
	"strings"
)

var list list.TodoList = list.TodoList{Date: "16/01/2025", }

var Commands = [3]string{	
	"pd", 
	"nd", 
	"add",
}

func splitStringIntoPrefix(input string) (string) {
	input = strings.TrimSpace(input)	


	before := strings.SplitN(input, " ", 2)
	
	return before[0]
}


func CheckisCommand(command string) bool { 
	if command == "" {
		return false
	}

	prefix := splitStringIntoPrefix(command)
	if prefix == "" {
		return false
	}

	for _, c := range Commands {
		if prefix == c {
			return true
		}
	}
	return false
}

func DoCommand(command string) {
	
	command = strings.TrimSpace(command)
	prefix := splitStringIntoPrefix(command)

	
	switch prefix {
	case "pd":
		print("previous")
	case "nd":
		print("next")
	case "add": 
		parts := strings.SplitN(command, "", 2)
		item := ""

		if len(parts) > 1 {
			item = parts[1]
		}
		fmt.Printf("added %s\n", item)
	}
}

func add(itemAdded string) { 
	
}