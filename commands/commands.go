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
	y := strconv.Itoa(Year)
	m := int(Month)
	mStr := strconv.Itoa(m)
	d := strconv.Itoa(Day)


	if (timeTraversed == 0) {
		return y + "-" + mStr + "-" + d
	}

	if (timeTraversed > 0) {
		newDay := Day + timeTraversed
		strNewDay := strconv.Itoa(newDay)
		newDate := y + "-" + mStr + "-" + (strNewDay)

		return newDate
	}
	

	return ""

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
