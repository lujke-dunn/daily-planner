package commands 

import "fmt"


type commandTitle struct {
	command string
	commandDesc string
}

func getCommand(ct *commandTitle) string {
	return ct.command
}



// map of command names
var commandName = map[commandTitle]int {
	{command: "pd", commandDesc: "Previous Day"}: 0,
	{command: "nd", commandDesc: "Next Day"}: 1,
	{command: "add", commandDesc: "Add Item To Current Days List"}: 2, 
	{command: "delete", commandDesc: "Delete Item at Index Specified"}: 3, 
	{command: "help", commandDesc: "Lists all commands"} : 4,
	
	
}
