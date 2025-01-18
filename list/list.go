package list


import (
  "fmt"
)

type TodoList struct {
	Date string
	ListItems []string
}


func (t *TodoList) AppendItems(items ...string) { 
	t.ListItems = append(t.ListItems, items...)
}

func (t *TodoList) PrintList() {
  for i, items := range t.ListItems {
    fmt.Printf("%d. %s\n", i+1, items)
  }
}
