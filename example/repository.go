package example

import "github.com/gilperopiola/go-mocks-architecture-poc/example/core"

// RepositoryI is the interface for the Repository, it holds all of its methods
type RepositoryI interface {
	GetUser(id int) (*core.User, error)
	IsUserValid(id int) bool
}

// Repository is the actual Repository. This would be used in a real life scenario
type Repository struct{}

// GetUser should be the actual implementation of the method. Used in real life
func (r *Repository) GetUser(id int) (*core.User, error) {
	return &core.User{}, nil
}

func (r *Repository) IsUserValid(id int) bool {
	return false
}
