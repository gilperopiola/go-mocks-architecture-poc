package example

import (
	"fmt"

	"github.com/gilperopiola/go-mocks-architecture-poc/example/core"
)

func main() {
	svc := Svc{Repository: &Repository{}}
	svc.Run()
}

// Here we have an App struct, with a RepositoryI inside of it.
// This RepositoryI can be either a Repository or a RepositoryMock, for testing.
// On main_test.go we test the IsUserValid method, needing to mock the Repository.GetUser(id) call.

type Svc struct {
	Repository RepositoryI
}

// IsUserValid is the method tested on main_test.go
func (s *Svc) GetUser(id int, checkIdentity bool) (*core.User, error) {
	user, err := s.Repository.GetUser(id)
	if err != nil {
		return &core.User{}, fmt.Errorf("GetUser: s.Repository.GetUser(id) error: %v", err)
	}

	if user.ID == 0 {
		return &core.User{}, fmt.Errorf("GetUser: user.ID == 0")
	}

	if checkIdentity && !s.Repository.IsUserValid(id) {
		return &core.User{}, fmt.Errorf("GetUser: !s.Repository.IsUserValid(id)")
	}

	return user, nil
}

func (s *Svc) Run() {
	fmt.Println("Service is running! (not really, it's just a POC)")
}
