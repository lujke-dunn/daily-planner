package commands 

import (
	"src/list"
	"fmt"
	"strings"
  "time"
	"strconv"
)


var currentTime = time.Now()

var Year = currentTime.Year()
var Month = currentTime.Month()
var Day = currentTime.Day()

func changeOrGetDate(Year int, Month time.Month , Day int, timeTraversed int) string {
		if timeTraversed == 0 {
			return formatDate(Year, Month, Day)
		}
		
		currentDate := time.Date(Year, Month, Day, 0, 0, 0, 0, time.UTC)

		newDate := currentDate.AddDate(0, 0, timeTraversed) // year, month, days ; moved 
		
		return formatDate(newDate.Year(), newDate.Month(), newDate.Day())
}

func formatDate(year int, month time.Month, day int) string {
	monthInt := int(month)

	monthStr := fmt.Sprintf("%02d", monthInt)
	dayStr := fmt.Sprintf("%02d", day)
	
	return fmt.Sprintf("%s-%s-%d", dayStr, monthStr, year)

}


var TodoList list.TodoList = list.TodoList{Date: changeOrGetDate(Year, Month, Day, 0), ListItems: []string{}}

func newTodoListDate(dateToChange string) {
	TodoList.ModifyDate(dateToChange)
}


var Commands = [5]string{	
  "date",
	"cd", 
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
    currentTime := TodoList.Date
    print(currentTime)

	case "cd":
		parts := strings.SplitN(command, " ", 2)
		daysMoved := ""

		if len(parts) > 1 {
			daysMoved = parts[1]
		}
		daysMovedInt, err := strconv.Atoi(daysMoved)


		if err != nil {
			print(err)
		} else {
			newTodoListDate(changeOrGetDate(Year, Month, Day, daysMovedInt))
			print(changeOrGetDate(Year, Month, Day, daysMovedInt))

		}		
	case "add": 
		parts := strings.SplitN(command, " ", 2)
		item := ""

		if len(parts) > 1 {
			item = parts[1]
      add(item)
      fmt.Print(TodoList.GetList())
		} else {
      fmt.Println("no item added")
    }
		case "list":
			fmt.Print(TodoList.GetList())
	}
}

func add(itemAdded string) { 
	TodoList.AppendItems(itemAdded)
  
}
