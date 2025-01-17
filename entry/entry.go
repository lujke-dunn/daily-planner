package entry

import (
	"bufio"
	"fmt"
	"io"
	"src/commands"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		text := scanner.Text()

		if (commands.CheckisCommand(text)) {
			commands.DoCommand(text)
		} else {
			giveError()
		}	
	}
}


func giveError() {
	fmt.Print("Invalid Command")
}
