package example

import (
	"fmt"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
)

// This is the first example of the go-mocks-architecture-poc.
// It's a little more complex than the one on the base folder, but it's simple still

func main() {
	svc := Service{Repository: &Repository{}}
	svc.Run()
}

// Our Service has a RepositoryI on it.
// This RepositoryI can be either a Repository or a RepositoryMock, either the real thing or the one we use on tests
type Service struct {
	Repository RepositoryI
}

// On example/main_test.go we test this method, with mocks for the s.Repository.GetUser(id) and s.Repository.IsUserValid(id) calls
func (s *Service) GetUser(id int, checkIdentity bool) (*core.User, error) {
	user, err := s.Repository.GetUser(id)
	if err != nil {
		return &core.User{}, errGettingUser
	}

	if user.ID == 0 {
		return &core.User{}, errInvalidID
	}

	if checkIdentity && !s.Repository.IsUserValid(id) {
		return &core.User{}, errInvalidUser
	}

	return user, nil
}

var (
	errGettingUser = fmt.Errorf("error getting user")
	errInvalidID   = fmt.Errorf("error: invalid id")
	errInvalidUser = fmt.Errorf("error: invalid user")
)

func (s *Service) Run() {
	fmt.Println("Service is running! (not really, it's just a POC)")
}
