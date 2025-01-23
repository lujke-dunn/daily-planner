package list


import (
  "fmt"
)

type TodoList struct {
	Date string
	ListItems []string
}

type TodoManager struct {
	Lists map[string]*TodoList
}

func CreateTodoManager() *TodoManager {
	return &TodoManager{
		Lists: make(map[string]*TodoList), 
	}
}

func (tm *TodoManager) GetOrCreateTodoList(date string) *TodoList {
	if list, exists := tm.Lists[date]; exists {
		return list
	}

	newList := &TodoList{
		Date: date, 
		ListItems: []string{},
	}
	tm.Lists[date] = newList
	return newList
}


func (t *TodoList) ModifyDate(newDate string) {
	t.Date = newDate
}


func (t *TodoList) DeleteItem(index int) error {
	if index < 0 || index >= len(t.ListItems) {
		return fmt.Errorf("index out of range")
	}
	t.ListItems = append(t.ListItems[:index], t.ListItems[+1:]...)
	return nil
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
