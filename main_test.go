package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppIsUserValid(t *testing.T) {
	tests := []struct {
		name      string
		argUserID int
		want      bool

		mockRepository func() *RepositoryMock
	}{
		{
			name:      "success",
			argUserID: 1,
			want:      true,

			mockRepository: func() *RepositoryMock {
				values := repositoryMockrepositoryGetUserOptions["default"]
				return setupRepositoryMockWithGetUser(values.UserID, values.Response, values.Error)
			},
		},
		{
			name:      "error_getting_user",
			argUserID: 0,
			want:      false,

			mockRepository: func() *RepositoryMock {
				values := repositoryMockrepositoryGetUserOptions["none"]
				return setupRepositoryMockWithGetUser(values.UserID, values.Response, values.Error)
			},
		},
		{
			name:      "error_invalid_id",
			argUserID: 0,
			want:      false,

			mockRepository: func() *RepositoryMock {
				values := repositoryMockrepositoryGetUserOptions["err_not_found"]
				return setupRepositoryMockWithGetUser(values.UserID, values.Response, values.Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			app := App{Repository: tt.mockRepository()}

			// Act
			got := app.IsUserValid(tt.argUserID)

			// Assert
			assert.Equal(t, tt.want, got)
		})
	}
}
