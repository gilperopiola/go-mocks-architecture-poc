package example

import (
	"testing"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
	"github.com/gilperopiola/go-mocks-architecture-poc/example/mocks"

	"github.com/stretchr/testify/assert"
)

// TestSvcGetUser tests a method that makes 2 calls to other methods, both belonging to the Repository.
// So, to test Svc.GetUser in isolation, we have to mock both Repository.GetUser and Repository.IsUserValid.

func TestSvcGetUser(t *testing.T) {

	type deps struct {
		repository func() *mocks.RepositoryMock
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
			out:  out{user: mocks.GetUserOptions["default"].Response, err: nil},

			deps: deps{repository: func() *mocks.RepositoryMock {
				return setupWithGetUserAndIsUserValid("default", "valid")
			}},
		},
		{
			name: "error_invalid_user",
			in:   in{userID: 1, checkIdentity: true},
			out:  out{user: mocks.GetUserOptions["none"].Response, err: errInvalidUser},

			deps: deps{repository: func() *mocks.RepositoryMock {
				return setupWithGetUserAndIsUserValid("default", "invalid")
			}},
		},
		{
			name: "error_invalid_id",
			in:   in{userID: 0, checkIdentity: true},
			out:  out{user: mocks.GetUserOptions["none"].Response, err: errInvalidID},

			deps: deps{repository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["none"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)
				return mock
			}},
		},
		{
			name: "error_getting_user",
			in:   in{userID: 0, checkIdentity: true},
			out:  out{user: mocks.GetUserOptions["err_not_found"].Response, err: errGettingUser},

			deps: deps{repository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["err_not_found"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)
				return mock
			}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// Prepare
			svc := Svc{Repository: tc.deps.repository()}

			// Act
			got, err := svc.GetUser(tc.in.userID, tc.in.checkIdentity)

			// Assert
			assert.Equal(t, tc.out.user, got)
			assert.Equal(t, tc.out.err, err)

			tc.deps.repository().AssertExpectations(t)
		})
	}
}

func setupWithGetUserAndIsUserValid(getUserOption, isUserValidOption string) *mocks.RepositoryMock {
	getUserValues := mocks.GetUserOptions[getUserOption]
	isUserValidValues := mocks.IsUserValidOptions[isUserValidOption]

	mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)
	return mocks.SetupWithIsUserValid(mock, isUserValidValues.UserID, isUserValidValues.Valid)
}
