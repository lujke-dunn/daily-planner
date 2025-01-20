package list


import (
  "fmt"
)

type TodoList struct {
	Date string
	ListItems []string
}


func (t *TodoList) ModifyDate(newDate string) {
	t.Date = newDate
}

func (t *TodoList) AppendItems(items ...string) { 
	t.ListItems = append(t.ListItems, items...)
}

func (t *TodoList) GetList() string {	
	var a string = "" 


  for i, items := range t.ListItems {
    a += fmt.Sprintf("%d. %s\n", i+1, items)
  }


	return a
}
