# Go Mocks Architecture - Proof of Concept

This is just a POC of a way to structure mocks and use them effectively on table-driven unit tests in Go. It focuses on maximizing simplicity when writing or reading tests, giving a name and predefining the different scenarios that the mock will encounter, thus allowing different tests to reuse those predefined options.

```bash
var mockMethodOptions = map[string]struct {
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

func TestSomething(t *testing.T) etc {
  values := mockMethodOptions["default"]
  mock := setupMock(values.Argument, values.Response)

  app := App{Repository: mock}
  assert.True(t, app.Test())
}
```

`What?` So the different options a mock method can behave in are defined alongside the mock, making them reusable for every test that needs to mock that method? `Just incredible! 🥳`

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
