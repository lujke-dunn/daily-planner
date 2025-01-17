package list


type TodoList struct {
	Date string
	ListItems []string
}


func (t *TodoList) AppendItems(items ...string) { 
	t.ListItems = append(t.ListItems, items...)
}
