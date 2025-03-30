package repository

import "fmt"

var ErrTodoNotFound = fmt.Errorf("todo entry not found")
var ErrListNotFound = fmt.Errorf("list entry not found")
