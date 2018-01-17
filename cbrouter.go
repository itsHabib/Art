package nap

import (
	"net/http"
)

// RouterFunc defines a callback router with optional content in the
// second parameter
type RouterFunc func(*http.Client, interface{})

// CBRouter maps status codes to specific callback functions
type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}
