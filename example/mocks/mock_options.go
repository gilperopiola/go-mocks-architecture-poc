package mocks

import (
	"fmt"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
)

var Options = struct {
	RepositoryGetUser     map[string]repositoryGetUserValues
	RepositoryIsUserValid map[string]repositoryIsUserValidValues
}{

	RepositoryGetUser: map[string]repositoryGetUserValues{
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
	},

	RepositoryIsUserValid: map[string]repositoryIsUserValidValues{
		"valid": {
			UserID: 1,
			Valid:  true,
		},
		"invalid": {
			UserID: 1,
			Valid:  false,
		},
	},
}
