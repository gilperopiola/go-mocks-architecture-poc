package main

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// RepositoryMock: implements the RepositoryI interface. Has mock methods
type RepositoryMock struct {
	*mock.Mock
}

// newRepositoryMock: creates a new *RepositoryMock
func newRepositoryMock() *RepositoryMock {
	return &RepositoryMock{&mock.Mock{}}
}

// GetUser: mock method. On tests, this will be called instead of the actual GetUser method
func (m *RepositoryMock) GetUser(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// repositoryMockGetUserScenarios: map with the different scenarios the mock will handle.
// Each scenario has the arguments and return values that the method call will use
var repositoryMockGetUserScenarios = map[string]struct {
	// Arguments
	UserID int

	// Return values
	Response User
	Error    error
}{
	"none": {
		UserID:   0,
		Response: User{},
		Error:    nil,
	},
	"default": {
		UserID:   1,
		Response: User{ID: 1, Username: "John"},
		Error:    nil,
	},
	"err_not_found": {
		UserID:   0,
		Response: User{},
		Error:    fmt.Errorf("user not found"),
	},
}

// setupRepositoryMockWithGetUser: creates the mock and calls the method
func setupRepositoryMockWithGetUser(userID int, response User, err error) *RepositoryMock {
	repositoryMock := newRepositoryMock()
	repositoryMock.On("GetUser", userID).Return(response, err).Once()
	return repositoryMock
}
