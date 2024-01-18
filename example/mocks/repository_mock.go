package mocks

import (
	"fmt"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	*mock.Mock
}

func newRepositoryMock() *RepositoryMock {
	return &RepositoryMock{&mock.Mock{}}
}

/* Get User Mock */

func (m *RepositoryMock) GetUser(id int) (*core.User, error) {
	args := m.Called(id)
	return args.Get(0).(*core.User), args.Error(1)
}

var GetUserOptions = map[string]struct {
	// Params
	UserID int
	// Return values
	Response *core.User
	Error    error
}{
	"none": {
		UserID:   0,
		Response: &core.User{},
		Error:    nil,
	},
	"default": {
		UserID:   1,
		Response: &core.User{ID: 1, Username: "John"},
		Error:    nil,
	},
	"err_not_found": {
		UserID:   0,
		Response: &core.User{},
		Error:    fmt.Errorf("user not found"),
	},
}

func SetupWithGetUser(userID int, response *core.User, err error) *RepositoryMock {
	repositoryMock := newRepositoryMock()
	repositoryMock.On("GetUser", userID).Return(response, err).Once()
	return repositoryMock
}

/* Is User Valid Mock */

func (m *RepositoryMock) IsUserValid(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

var IsUserValidOptions = map[string]struct {
	// Params
	UserID int
	// Return values
	Valid bool
}{
	"valid": {
		UserID: 1,
		Valid:  true,
	},
	"invalid": {
		UserID: 1,
		Valid:  false,
	},
}

func SetupWithIsUserValid(mock *RepositoryMock, userID int, response bool) *RepositoryMock {
	mock.On("IsUserValid", userID).Return(response).Once()
	return mock
}
