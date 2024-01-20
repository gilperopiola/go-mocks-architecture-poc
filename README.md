# Go Mocks Architecture ⭐ 

**`⬤ Welcome!` ~~~ `GOLANG 🐹 · UNIT TESTS 🔍 · MOCKS 🎭`**

🐋 This is an easy pattern designed for **SIMPLE** · **READABLE** · **UNDERSTANDABLE** · **REUSABLE CODE** 🐋

****

### Quickest tour ⚡

*`We have a method that we need to test:`*


```go
func (a *App) GetUser(id int) *User {
	if user, err := a.Repository.GetUser(id); err == nil {
		return user
	}

	return nil
}
```
*`a.Repository.GetUser(id) is the actual call that we are mocking. For each scenario (success, user not found, etc.) we define the arguments that call is going to take, and the values it will return:`*

```go
var getUserScenarios = map[string]struct {
	UserID int
	ResponseUser *User
}{
	"none": {
		UserID: 0,
		ResponseUser: nil,
	},
	"default": {
		UserID: 1,
		ResponseUser: &User{},
	},
}

// --- AND THEN JUST COPY-PASTE-CHANGE THIS FOR EACH METHOD:

func setupMock(userID int, responseUser *User) *RepositoryMock {
	mock := newRepositoryMock()
	mock.On("GetUser", userID).Return(responseUser).Once()
	return mock
}

// --- TO PUT IT ALL IN PLACE:

func TestGetUserSuccess(t *testing.T) {
	values := getUserScenarios["default"]
	mock := setupMock(values.UserID, values.ResponseUser)
	app := &App{Repository: mock}
	assert.Equals(t, app.GetUser(values.Argument), values.ResponseUser)
}

func TestGetUserError(t *testing.T) {
	values := getUserScenarios["none"]
	mock := setupMock(values.UserID, values.ResponseUser)
	app := &App{Repository: mock}
	assert.Equals(t, app.GetUser(values.Argument), values.ResponseUser)
}

// --- OF COURSE THIS CAN BE APPLIED TO TABLE-DRIVEN TESTS OR WHATEVER FORMAT YOU USE.
```



Unit tests are often left unpolished, messy, dirty. 

And if you can't read or understand what the test cases are doing, you won't be able to fix them when they break. 

So you spend some time failing to fully grasp the intent of each case before commenting it away and starting anew. Happens.

So I've decided to spend a tiny bit of time on the short run (setting up mock scenarios and writing some simple almost-copy-paste functions) to:

 - Improve tests simplicity and readability, understandability.
 - Be able to re-use mock scenarios in different tests.
 - Gain back some of the hours lost to aimless coding.

This is just a POC of a way to structure mocks and use them effectively on table-driven unit tests in Go. It focuses on maximizing simplicity when writing or reading tests, giving a name and predefining the different scenarios that the mock will encounter, thus allowing different tests to reuse those predefined options.

```bash
var mockMethodScenarios = map[string]struct {
  Argument int
  Response *User
}{
	"none": {
		Argument: 0,
		Response: nil,
	},
	"default": {
		Argument: 1,
		Response: &User{},
	},
}

func setupMock(argument int, response *User) *RepositoryMock {
	mock := newRepositoryMock()
	mock.On("MyMethod", argument).Return(response).Once()
	return mock
}

func TestSomething(t *testing.T) {
	values := mockMethodScenarios["default"]
	mock := setupMock(values.Argument, values.Response)

	app := App{Repository: mock}
	assert.True(t, app.Test())
}
```

`What!?😱` So the different ways a mock method can behave in are defined alongside the mock, making them reusable for every test that needs to mock that method? `Just incredible!!! 🥳`.

## How do I use it? 🤔

```bash
func TestMethod(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		mock      func() *RepositoryMock
	}{
		{
			name:      "success",
			want:      true,
			mock: func() *RepositoryMock {
				values := mockMethodOptions["default"]
				return setupMock(values.Argument, values.Response)
			},
		},
		{
			name:      "failure",
			want:      false,
			mock: func() *RepositoryMock {
				values := mockMethodOptions["none"]
				return setupMock(values.Argument, values.Response)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := App{Repository: tt.mock()}
			got := app.Method()
			assert.Equal(t, tt.want, got)
		})
	}
}
```

For each test case we use a different key for the `mockMethodOptions` map, like `["default"]` or `["none"]`. 

Upcoming tests can choose to reuse the existing options, modify them (taking care of the other tests that use said option) or add new ones to the map.

## Now what?

On the `repository_mock.go` and `main_test.go` files you will find a simple example of how to implement this architecture.

Hopefully I will finish the example on `/example`, tailored for more complex projects where the mocks are on a different package and method calls are just `w0nky~`.
