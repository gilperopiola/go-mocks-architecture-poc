package example

import (
	"fmt"
	"testing"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
	"github.com/gilperopiola/go-mocks-architecture-poc/example/mocks"

	"github.com/stretchr/testify/assert"
)

func TestSvcGetUser(t *testing.T) {
	tests := []struct {
		name             string
		argUserID        int
		argCheckIdentity bool
		want             *core.User
		wantErr          error

		mockRepository func() *mocks.RepositoryMock
	}{
		{
			name:             "success",
			argUserID:        1,
			argCheckIdentity: true,
			want:             mocks.GetUserOptions["default"].Response,
			wantErr:          nil,

			mockRepository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["default"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)

				isUserValidValues := mocks.IsUserValidOptions["valid"]
				mock = mocks.SetupWithIsUserValid(mock, isUserValidValues.UserID, isUserValidValues.Valid)

				return mock
			},
		},
		{
			name:             "error_invalid_user",
			argUserID:        1,
			argCheckIdentity: true,
			want:             mocks.GetUserOptions["none"].Response,
			wantErr:          fmt.Errorf("GetUser: !s.Repository.IsUserValid(id)"),

			mockRepository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["default"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)

				isUserValidValues := mocks.IsUserValidOptions["invalid"]
				mock = mocks.SetupWithIsUserValid(mock, isUserValidValues.UserID, isUserValidValues.Valid)

				return mock
			},
		},
		{
			name:             "error_invalid_id",
			argUserID:        0,
			argCheckIdentity: true,
			want:             mocks.GetUserOptions["none"].Response,
			wantErr:          fmt.Errorf("GetUser: user.ID == 0"),

			mockRepository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["none"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)
				return mock
			},
		},
		{
			name:             "error_getting_user",
			argUserID:        0,
			argCheckIdentity: true,
			want:             mocks.GetUserOptions["err_not_found"].Response,
			wantErr:          fmt.Errorf("GetUser: s.Repository.GetUser(id) error: %v", fmt.Errorf("user not found")),

			mockRepository: func() *mocks.RepositoryMock {
				getUserValues := mocks.GetUserOptions["err_not_found"]
				mock := mocks.SetupWithGetUser(getUserValues.UserID, getUserValues.Response, getUserValues.Error)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			svc := Svc{Repository: tt.mockRepository()}

			// Act
			got, err := svc.GetUser(tt.argUserID, tt.argCheckIdentity)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
