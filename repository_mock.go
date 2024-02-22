package main

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// repositoryMock is a mock of Repository interface.
type repositoryMock struct {
	*mock.Mock
}

// NewRepositoryMock returns a new mock of the Repository interface.
func newRepositoryMock() *repositoryMock {
	return &repositoryMock{&mock.Mock{}}
}

// GetUser: mock method. On tests, this will be called instead of the actual GetUser method.
func (m *repositoryMock) GetUser(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// repositoryMockGetUserOptions: map with the different options the mock will handle.
// Each option has the arguments and return values that the method call will use.
var repositoryMockGetUserOptions = map[string]struct {
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
		Response: User{ID: 1, Username: "gilperopiola"},
		Error:    nil,
	},
	"err_not_found": {
		UserID:   0,
		Response: User{},
		Error:    fmt.Errorf("user not found"),
	},
}

// setupRepositoryMockWithGetUser: creates the mock and calls the method.
func setupRepositoryMockWithGetUser(userID int, response User, err error) *repositoryMock {
	repositoryMock := newRepositoryMock()
	repositoryMock.On("GetUser", userID).Return(response, err).Once()
	return repositoryMock
}
