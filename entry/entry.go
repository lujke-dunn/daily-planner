package entry

import (
	"bufio"
	"fmt"
	"io"
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
		io.WriteString(out, text)		
	}
}