package main

import (
  "src/entry"
	"os"
)


// entry 
func main() {
	entry.Start(os.Stdin, os.Stdout) 
}
