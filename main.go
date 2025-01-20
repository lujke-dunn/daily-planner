package main

import (
  "src/entry"
	"os"
)



func main() {
	entry.Start(os.Stdin, os.Stdout) 
}
