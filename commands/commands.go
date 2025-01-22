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
var TodoManager = list.CreateTodoManager()
var CurrentDate = changeOrGetDate(Year, Month, Day, 0)
var CurrentList = TodoManager.GetOrCreateTodoList(CurrentDate)

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


func newTodoListDate(dateToChange string) {
	CurrentDate = dateToChange
	CurrentList = TodoManager.GetOrCreateTodoList(CurrentDate)

}


var Commands = [5]string{	
	"help",
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
      fmt.Print(CurrentList.GetList())
		} else {
      fmt.Println("no item added")
    }
		case "list":
			fmt.Print(CurrentList.GetList())
		case "help":
			fmt.Print("Try these commands:\n")
			fmt.Print("add ___: adds a item to the list, takes a string as an argument\n")
			fmt.Print("cd ___: changes the current viewing date. takes a int as an argument positive integers go into the future by x days, negative integers go into the past by -x days. Note: cd is relative.\n")
			fmt.Print("help: displays list of commands")
	}
}

func add(itemAdded string) { 
	CurrentList.AppendItems(itemAdded)
	fmt.Print("")
  
}
