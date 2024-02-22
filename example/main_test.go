package example

import (
	"testing"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
	"github.com/gilperopiola/go-mocks-architecture-poc/example/mocks"

	"github.com/stretchr/testify/assert"
)

// TestServiceGetUser tests a Service method that makes 2 calls to external Repository methods.
// So, to test Svc.GetUser in isolation, we have to mock both Repository.GetUser and Repository.IsUserValid calls.
func TestServiceGetUser(t *testing.T) {

	// Service dependencies
	type deps struct {
		repository *mocks.RepositoryMock
	}

	// Service.GetUser call parameters
	type in struct {
		userID        int
		checkIdentity bool
	}

	// Service.GetUser call expected output
	type out struct {
		user *core.User
		err  error
	}

	// Test cases
	TCs := []struct {
		name string
		deps deps
		in   in
		out  out
	}{
		{
			name: "success",
			in:   in{userID: 1, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["default"].Response, err: nil},
			deps: deps{repository: setupRepositoryWithGetUserAndIsUserValid("default", "valid")},
		},
		{
			name: "error_invalid_user",
			in:   in{userID: 1, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["none"].Response, err: errInvalidUser},
			deps: deps{repository: setupRepositoryWithGetUserAndIsUserValid("default", "invalid")},
		},
		{
			name: "error_invalid_id",
			in:   in{userID: 0, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["none"].Response, err: errInvalidID},
			deps: deps{repository: mocks.SetupRepositoryWithGetUser("none")},
		},
		{
			name: "error_getting_user",
			in:   in{userID: 0, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["err_not_found"].Response, err: errFailedToGetUser},
			deps: deps{repository: mocks.SetupRepositoryWithGetUser("err_not_found")},
		},
	}
	for _, tc := range TCs {
		t.Run(tc.name, func(t *testing.T) {

			// Prepare
			deps, in, out := tc.deps, tc.in, tc.out
			service := Service{Repository: deps.repository}

			// Act
			got, err := service.GetUser(in.userID, in.checkIdentity)

			// Assert
			assert.Equal(t, out.user, got)
			assert.Equal(t, out.err, err)

			service.Repository.(*mocks.RepositoryMock).AssertExpectations(t)
		})
	}
}

// setupRepositoryWithGetUserAndIsUserValid is used when both Repository methods are called, allowing an option for each one
func setupRepositoryWithGetUserAndIsUserValid(getUserOption, isUserValidOption string) *mocks.RepositoryMock {
	mock := mocks.SetupRepositoryWithGetUser(getUserOption)
	mock = mocks.SetupRepositoryWithIsUserValid(mock, isUserValidOption)
	return mock
}
