package example

import (
	"fmt"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
)

// Hey :)
// This is the first example of the Go Mocks Architecture POC.

var (
	errFailedToGetUser = fmt.Errorf("error: failed to get user")
	errInvalidID       = fmt.Errorf("error: invalid id")
	errInvalidUser     = fmt.Errorf("error: invalid user")
)

// This is the method we actually test on example/main_test.go.
// Both calls to s.Repository.GetUser(id) and s.Repository.IsUserValid(id) are mocked.
// Both of those are calls to the Repository, the interface that we mock.
// We test the different scenarios that the Service can run into, and those depend on the result of the calls to the Repository.
func (s *Service) GetUser(id int, checkIdentity bool) (*core.User, error) {

	// Scenario 1: GetUser returns an error
	user, err := s.Repository.GetUser(id)
	if err != nil {
		return &core.User{}, errFailedToGetUser
	}

	// Scenario 2: GetUser returns an invalid ID
	if user.ID == 0 {
		return &core.User{}, errInvalidID
	}

	// Scenario 3: GetUser returns a valid user, but checkIdentity is true and IsUserValid returns false
	if checkIdentity && !s.Repository.IsUserValid(id) {
		return &core.User{}, errInvalidUser
	}

	// Scenario 4: Success
	return user, nil
}

// init is used instead of main because we are on the example package.
// We simulate a simple API.
func init() {
	svc := Service{Repository: NewRepository()}
	svc.Run()
}

// Our Service has a Repository inside of it.
// This Repository can be either a repository or a repositoryMock, either the real thing or the one we use on tests.
type Service struct {
	Repository Repository
}

// Run is here to simulate the service running.
func (s *Service) Run() {
	fmt.Println("Service is running! (not really, it's just a POC)")
}
