package example

import (
	"testing"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
	"github.com/gilperopiola/go-mocks-architecture-poc/example/mocks"

	"github.com/stretchr/testify/assert"
)

// TestSvcGetUser tests a method that makes 2 calls to external Repository methods.
// So, to test Svc.GetUser in isolation, we have to mock both Repository.GetUser and Repository.IsUserValid

func TestSvcGetUser(t *testing.T) {

	type deps struct {
		repository *mocks.RepositoryMock
	}

	type in struct {
		userID        int
		checkIdentity bool
	}

	type out struct {
		user *core.User
		err  error
	}

	tests := []struct {
		name string
		deps deps
		in   in
		out  out
	}{
		{
			name: "success",
			in:   in{userID: 1, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["default"].Response, err: nil},
			deps: deps{repository: setupWithGetUserAndIsUserValid("default", "valid")},
		},
		{
			name: "error_invalid_user",
			in:   in{userID: 1, checkIdentity: true},
			out:  out{user: mocks.Options.RepositoryGetUser["none"].Response, err: errInvalidUser},
			deps: deps{repository: setupWithGetUserAndIsUserValid("default", "invalid")},
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
			out:  out{user: mocks.Options.RepositoryGetUser["err_not_found"].Response, err: errGettingUser},
			deps: deps{repository: mocks.SetupRepositoryWithGetUser("err_not_found")},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// Prepare
			svc := Svc{Repository: tc.deps.repository}

			// Act
			got, err := svc.GetUser(tc.in.userID, tc.in.checkIdentity)

			// Assert
			assert.Equal(t, tc.out.user, got)
			assert.Equal(t, tc.out.err, err)

			svc.Repository.(*mocks.RepositoryMock).AssertExpectations(t)
		})
	}
}

// setupWithGetUserAndIsUserValid is used when both Repository methods are called, allowing an option for each one
func setupWithGetUserAndIsUserValid(getUserOption, isUserValidOption string) *mocks.RepositoryMock {
	mock := mocks.SetupRepositoryWithGetUser(getUserOption)
	mock = mocks.SetupRepositoryWithIsUserValid(mock, isUserValidOption)
	return mock
}
