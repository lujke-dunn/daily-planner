package entry

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"os"
	"src/commands"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	
	fmt.Println("Hey %s! Welcome to the daily planner")

	scanner := bufio.NewScanner(in)
	fmt.Print(PROMPT)
	for scanner.Scan() {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := scanner.Text()

		if (commands.CheckisCommand(text)) {
			commands.DoCommand(text)
			clearTerminal()
			fmt.Print(DisplayRelevantInformation())
		  fmt.Print(PROMPT)

		} else {
			giveError()
		  fmt.Print(PROMPT)
		}	
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayRelevantInformation() string {
	
	date := commands.TodoList.Date
	list := commands.TodoList.GetList()


	ui := "========================" + date + "=======================\n"
	ui += list
	return ui 
}


func giveError() {
	fmt.Print("Invalid Command")
}
