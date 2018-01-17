package nap

// RestResource is the struct containing the REST endpoints
// corresponding methods, and corresponding routers for that resource
type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}
