package main

// RepositoryI is the interface for the Repository, it holds all of its methods
type RepositoryI interface {
	GetUser(id int) (User, error)
}

// Repository is the actual Repository. This would be used in a real life scenario
type Repository struct{}

// GetUser should be the actual implementation of the method. Used in real life
func (r *Repository) GetUser(id int) (User, error) {
	return User{}, nil
}

// User is the data structure used across the API
type User struct {
	ID       int
	Username string
}
