// Code generated by goa v3.7.3, DO NOT EDIT.
//
// HTTP request path constructors for the todo service.
//
// Command:
// $ goa gen todo/design

package client

import (
	"fmt"
)

// HelloTodoPath returns the URL path to the todo service hello HTTP endpoint.
func HelloTodoPath(name string) string {
	return fmt.Sprintf("/hello/%v", name)
}

// ShowTodoPath returns the URL path to the todo service show HTTP endpoint.
func ShowTodoPath(id int) string {
	return fmt.Sprintf("/user/%v", id)
}
