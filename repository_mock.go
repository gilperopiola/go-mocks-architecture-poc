package main

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

// RepositoryMock holds mock testing methods for the Repository
type RepositoryMock struct {
	*mock.Mock
}

func newRepositoryMock() *RepositoryMock {
	return &RepositoryMock{&mock.Mock{}}
}

// GetUser is a mock method. When the MockRepository is used on a test, this method
// will be called instead of the actual Repository's GetUser one
func (m *RepositoryMock) GetUser(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// repositoryMockrepositoryGetUserOptions has all the different options that the mock will handle.
// Each case has the parameters and return values that the method call will use
var repositoryMockrepositoryGetUserOptions = map[string]struct {
	// Params
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

// setupRepositoryMockWithGetUser is a helper, just creates the mock and places the call to the method
func setupRepositoryMockWithGetUser(userID int, response User, err error) *RepositoryMock {
	repositoryMock := newRepositoryMock()
	repositoryMock.On("GetUser", userID).Return(response, err).Once()
	return repositoryMock
}
