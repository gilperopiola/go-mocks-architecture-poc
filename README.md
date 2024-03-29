# Go Mocks Architecture ⭐ 

**`⬤ Welcome!` ~~~ `GOLANG 🐹 · UNIT TESTS 🔍 · MOCKS 🎭`**

🐋 This is an easy 'pattern' designed for **SIMPLE** · **READABLE** · **REUSABLE CODE** 🐋

🤔 Need to mock a `.GetUser(id)` call?

	Just define all different possible scenarios: success, success_admin, err_not_found, err_unknown!

And set the **input and return values** for each scenario **:)**.

----
### Quickest tour⚡

_**·** We wanna unit test this method:_

```go
func (a *App) DeleteUser(id int) *User {
	if user, err := a.Repository.GetUser(id); err == nil {
		// ... code to delete the user
		return user
	}

	return nil
}
```
_**·** So `a.Repository.GetUser(id)` is the actual call that we will mock. For each option (`success` or `user not found`) we define the **arguments** and **return values** of the method call:_

```go
var getUserMockoptions = map[string]struct {
	UserID       int   // argument
	UserToReturn *User // return value
}{
	"success": {
		UserID: 1,
		UserToReturn: &User{Name: "John"},
	},
	"user_not_found": {
		UserID: 0,
		UserToReturn: nil,
	},
}
```

_**·** All in all, you set up the mock with the values of your desired option:_

```go
func TestDeleteUserSuccess(t *testing.T) {
	// get the argument and return value for the call to Repository.GetUser
	values := getUserMockoptions["success"]

	// set up the mock using the option's predefined values
	app := &App{Repository: setupMockWithGetUser(values.UserID, values.UserToReturn)}

	got := app.DeleteUser(values.UserID)
	
	// assert results
	assert.Equals(t, values.UserToReturn, got)
	app.Repository.AssertExpectations(t)
}
```

```go
// helper function
func setupMockWithGetUser(userID int, userToReturn *User) *RepositoryMock {
	mock := &RepositoryMock{&mock.Mock{}}
	mock.On("GetUser", userID).Return(userToReturn).Once()
	return mock
}
```
----

### But is it any good? 🧐 

Well...

Simple code     | Readable code             | Understandable code      | 
---             | ---                       | ---                      |
Easy to pick up | Easy to master            | Easy to maintain         |
Follows DRY     | Simple & complex examples | Re-use mock options    |
Adaptable       | Excellent documentation   | **@gilperopiola** 🔥     |

### 🐿️

✅ **Works with any kind of test setup ;)**\
✅ **Supports TDD**\
✅ **Supports Table Driven Tests**

----
### Production ready real life example 📦

**Table driven tests**. The function to test has a call to `.Repository.GetUser(id int) *User` so we mock 2 different options: _one gets the user successfully and the other doesn't_.

```go
func TestIsValid(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		mock      func() *RepositoryMock
	}{
		{
			name: "valid",
			want: true,
			mock: func() *RepositoryMock {
				values := mockGetUserOptions["default"]
				return setupMockWithGetUser(values.UserID, values.UserToReturn)
			},
		},
		{
			name: "invalid",
			want: false,
			mock: func() *RepositoryMock {
				values := mockGetUserOptions["none"]
				return setupMockWithGetUser(values.UserID, values.UserToReturn)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := App{Repository: tt.mock()}
			got := app.IsValid()
			assert.Equal(t, tt.want, got)
		})
	}
}
```

For each test case we use a different key for the `mockGetUserOptions` map, like `["default"]` or `["none"]`. 

#### Upcoming tests can reuse the existing options, modify them (carefully😅) or add new ones to the map.
----
### Now what? 🐿️

On the `repository_mock.go` and `main_test.go` files you will find a simple example of how to implement this architecture.

On `/example` we have a more complex _example_, the mocks are on a different package and multiple mock methods are called.

----
### Motivation 🚀

	· Unit tests are often left unpolished, messy, dirty. 

	· And if you can't read or understand what the test cases are doing, you won't be able to fix them when they break. 

	· So you spend some time failing to fully grasp the intent of each case before commenting it away and starting anew. Happens.

	· So why not spend a tiny bit of time on the short run (setting up mock options and copy-pasting functions) to:

	· 	Improve tests simplicity and readability, understandability.
	· 	Be able to re-use mock options in different tests.
	· 	Gain back some of the hours lost to aimless coding.

	:)
