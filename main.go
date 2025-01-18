package main

import (
	"fmt"
  "os"
  "os/user"
  "src/entry"
)



func main() {
	user, err := user.Current()
  if err != nil {
		panic(err)
	}
	fmt.Printf("Hey %s! Welcome to the daily planner\n", user.Username)

	entry.Start(os.Stdin, os.Stdout)
 
  /*commands.DoCommand("add buy milk")
  commands.DoCommand("add write code")
  commands.DoCommand("list")
  */
  
}
