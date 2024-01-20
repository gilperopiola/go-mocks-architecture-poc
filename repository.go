package main

// RepositoryI: interface for the Repository, it has all of its methods
type RepositoryI interface {
	GetUser(id int) (User, error)
}

// Repository: the actual Repository. This would be used in a real life scenario
type Repository struct{}

// GetUser: the actual implementation of the method. Used in real life
func (r *Repository) GetUser(id int) (User, error) {
	return User{}, nil
}

// User: data structure used across the API
type User struct {
	ID       int
	Username string
}
