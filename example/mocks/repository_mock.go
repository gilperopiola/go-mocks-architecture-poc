package mocks

import (
	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"

	"github.com/stretchr/testify/mock"
)

// RepositoryMock is a mock of Repository interface.
type RepositoryMock struct {
	*mock.Mock
}

// NewRepositoryMock returns a new mock of the Repository interface.
func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{&mock.Mock{}}
}

/* ----------------------------------- */
/*          - Get User Mock -          */
/* ----------------------------------- */

func (m *RepositoryMock) GetUser(id int) (*core.User, error) {
	args := m.Called(id)
	return args.Get(0).(*core.User), args.Error(1)
}

type repositoryGetUserValues struct {
	UserID int

	Response *core.User
	Error    error
}

func SetupRepositoryWithGetUser(option string) *RepositoryMock {
	optionValues := Options.RepositoryGetUser[option]
	repositoryMock := NewRepositoryMock()
	repositoryMock.On("GetUser", optionValues.UserID).Return(optionValues.Response, optionValues.Error).Once()
	return repositoryMock
}

/* ----------------------------------- */
/*       - Is User Valid Mock -        */
/* ----------------------------------- */

func (m *RepositoryMock) IsUserValid(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

type repositoryIsUserValidValues struct {
	UserID int

	Valid bool
}

// SetupRepositoryWithIsUserValid takes a mock as parameter because a previous mocked call to the GetUser method is needed.
func SetupRepositoryWithIsUserValid(mock *RepositoryMock, option string) *RepositoryMock {
	optionValues := Options.RepositoryIsUserValid[option]
	mock.On("IsUserValid", optionValues.UserID).Return(optionValues.Valid).Once()
	return mock
}
