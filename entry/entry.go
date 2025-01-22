package entry

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"os"
	"src/commands"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	
	fmt.Println("Hey, Welcome to the daily planner")
	fmt.Print(DisplayRelevantInformation())

	scanner := bufio.NewScanner(in)
	fmt.Print(PROMPT)
	for scanner.Scan() {
		text := scanner.Text()

		if (commands.CheckisCommand(text)) {
			commands.DoCommand(text)
			clearTerminal()
			fmt.Print(DisplayRelevantInformation())
		  fmt.Print(PROMPT)

		} else {
			giveError()
			fmt.Print("\n"+PROMPT)
		}	
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayRelevantInformation() string {
	
	date := commands.CurrentList.Date
	list := commands.CurrentList.GetList()


	ui := "========================" + date + "=======================\n"
	ui += list
	return ui 
}


func giveError() {
	fmt.Print("Invalid Command\n")
	commands.DoCommand("help")
}
