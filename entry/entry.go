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
	return "do command"
}

func giveError() {
	fmt.Print("Invalid Command")
}
