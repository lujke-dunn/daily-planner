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

// creates a new todoManager which holds the list of lists 
func CreateTodoManager() *TodoManager {
	return &TodoManager{
		Lists: make(map[string]*TodoList), 
	}
}


// checks if a list on inputted date exists, else returns a new list with an empty list 
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

// date setter
func (t *TodoList) ModifyDate(newDate string) {
	t.Date = newDate
}

// check if the items index is positive and is within bounds and then remove the item
func (t *TodoList) DeleteItem(index int) error {
	if index < 0 || index >= len(t.ListItems) {
		return fmt.Errorf("index out of range")
	}
	t.ListItems = append(t.ListItems[:index], t.ListItems[+1:]...)
	return nil
}

// add item/s to the list
func (t *TodoList) AppendItems(items ...string) { 
	t.ListItems = append(t.ListItems, items...)
}


// returns all items on the list 
func (t *TodoList) GetList() string {	
	var a string = "" 


  for i, items := range t.ListItems {
    a += fmt.Sprintf("%d. %s\n", i+1, items)
  }


	return a
}
