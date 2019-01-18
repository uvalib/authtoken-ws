package dao

// our storage interface
type Storage interface {
	GetActiveTokens() (TokenPermissions, error)
	Destroy() error
}

// Permissions -- a permission associated with a token
type Permissions struct {
	Whom string
	What string
}

// TokenPermissions -- a map of tokens and their associated permissions
type TokenPermissions map[string]Permissions

// our singleton store
var Store Storage

// our factory
func NewDatastore( ) error {
	var err error
	// mock implementation here
	Store, err = newDBStore()
	return err
}

//
// end of file
//
