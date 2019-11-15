package interfaces

import "net/http"

// Context is Interface for Controller
type Context interface {
	Request() http.Request
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
