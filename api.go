package nap

// API is used to map strings to specific REST resources
// Simplifies making requests by using a simple string name
// to make requests
type API struct {
	BaseURL       string
	Resources     map[string]RestResource
	DefaultRouter *CBRouter
	Client        *Client
}
