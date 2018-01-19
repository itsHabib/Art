package art

// AuthToken defines a token used for auth
type AuthToken struct {
	Token string
}

// AuthBasic defines a Basic Auth structure
type AuthBasic struct {
	Username string
	Password string
}

// Authentication interface allows both implementations of basic auth
// && auth token
type Authentication interface {
	AuthorizationHeader() string
}
