package commands 

import (
	"src/list"
	"fmt"
	"strings"
  "time"
)

var TodoList list.TodoList = list.TodoList{Date: "16/01/2025", ListItems: []string{}}

var Commands = [5]string{	
  "date",
	"pd", 
	"nd", 
	"add",
  "list", 
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
  case "date": 
    currentTime := time.Now()
    print(currentTime.Year() - 1)

	case "pd":
		print("previous")
	case "nd":
		print("next")
	case "add": 
		parts := strings.SplitN(command, "", 2)
		item := ""

		if len(parts) > 1 {
			item = parts[1]
      add(item)
      TodoList.PrintList()
		} else {
      fmt.Println("no item added")
    }
  case "list":
      TodoList.PrintList()
	}
}

func add(itemAdded string) { 
	TodoList.AppendItems(itemAdded)
  
}
