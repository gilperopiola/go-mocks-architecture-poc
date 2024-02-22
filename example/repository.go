package example

import "github.com/gilperopiola/go-mocks-architecture-poc/example/core"

// Repository is the interface for the repository layer of the architecture.
// It should handle all the external communication and data access.
// This is the contract that the repository must fulfill.
type Repository interface {
	GetUser(id int) (*core.User, error)
	IsUserValid(id int) bool
}

// repository is the actual implementation of the Repository interface.
type repository struct{}

// NewRepository returns a new Repository.
func NewRepository() Repository {
	return &repository{}
}

// GetUser should be the actual implementation of the method.
func (r *repository) GetUser(id int) (*core.User, error) {

	// ... some code to get the user from the database
	// ... error handling

	return &core.User{ID: id, Username: "gilperopiola"}, nil
}

// IsUserValid should be the actual implementation of the method.
func (r *repository) IsUserValid(id int) bool {

	// ... some code to get the user from the database
	// ... check user validity

	return id > 0
}
