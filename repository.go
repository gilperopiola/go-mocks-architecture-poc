package main

// Repository is the interface for the repository layer of the architecture.
// It should handle all the external communication and data access.
// This is the contract that the repository must fulfill.
type Repository interface {
	GetUser(id int) (User, error)
}

// repository is the actual implementation of the Repository interface.
type repository struct{}

// NewRepository returns a new Repository.
func NewRepository() Repository {
	return &repository{}
}

// GetUser should be the actual implementation of the method.
func (r *repository) GetUser(id int) (User, error) {

	// ... some code to get the user from the database
	// ... error handling

	return User{ID: id, Username: "gilperopiola"}, nil
}

// User is the data structure used across the example API.
type User struct {
	ID       int
	Username string
}
