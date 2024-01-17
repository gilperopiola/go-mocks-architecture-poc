package main

import (
	"fmt"
)

func main() {
	app := App{Repository: &Repository{}}
	app.Run()
}

// Here we have an App struct, with a RepositoryI inside of it.
// This RepositoryI can be either a Repository or a RepositoryMock, for testing.
// On main_test.go we test the IsUserValid method, needing to mock the Repository.GetUser(id) call.

type App struct {
	Repository RepositoryI
}

// IsUserValid is the method tested on main_test.go
func (a *App) IsUserValid(id int) bool {
	user, err := a.Repository.GetUser(id)
	if err != nil {
		return false // error_getting_user
	}

	if user.ID == 0 {
		return false // error_invalid_id
	}

	return true // success
}

func (a *App) Run() {
	fmt.Println("App is running! (not really, it's just a POC)")
}
