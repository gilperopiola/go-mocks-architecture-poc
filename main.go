package main

import (
	"fmt"
)

// Hey :)
// Welcome to Go Mocks Architecture POC.

func main() {
	app := App{Repository: &repository{}}
	app.Run()
}

// App has a Repository inside of it.
// This Repository can be either a repository or a repositoryMock, for testing.
// On main_test.go we test the IsUserValid method, needing to mock the Repository.GetUser(id) call.
type App struct {
	Repository Repository
}

// IsUserValid is the method tested on main_test.go
func (a *App) IsUserValid(id int) bool {

	// Scenario 1: GetUser returns an error
	user, err := a.Repository.GetUser(id)
	if err != nil {
		return false
	}

	// Scenario 2: GetUser returns an invalid ID
	if user.ID == 0 {
		return false
	}

	// Scenario 3: Success
	return true
}

// Run is here to simulate the app running.
func (a *App) Run() {
	fmt.Println("App is running! (not really, it's just a POC)")
}
