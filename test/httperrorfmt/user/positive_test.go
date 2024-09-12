package user

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	"app/test/setup"
	"reflect"
	"testing"
)

func TestPositiveUser(t *testing.T) {
	app := setup.Run()
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 1

		// act
		users, err := app.HttpErrorFmts.User.GetAll(1, 0, 50)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}
		got := len(*users)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := m.User{Id: 1, FirstName: "John", LastName: "Doe", Username: "john_doe", Email: "john@doe.com"}

		// act
		got, err := app.HttpErrorFmts.User.GetOne(1, 1)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Add", func(t *testing.T) {
		// arrange
		got := m.NewUser("James", "Doe", "james_doe", "james@doe.com", "password1")
		expect := m.User{Id: 2, FirstName: "James", LastName: "Doe", Username: "james_doe", Email: "james@doe.com"}

		// act
		err := app.HttpErrorFmts.User.Add(1, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		got := &m.EditUser{FirstName: "John1", LastName: "Doe1", Username: "john_doe", Email: "john@doe.com"}
		expect := m.EditUser{Id: 1, FirstName: "John1", LastName: "Doe1", Username: "john_doe", Email: "john@doe.com"}

		// act
		err := app.HttpErrorFmts.User.Edit(1, 1, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expect := e.NewErrRepoNotFound("user id", "2")

		// act
		err := app.HttpErrorFmts.User.Remove(1, 2)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}
		_, got := app.HttpErrorFmts.User.GetOne(1, 2)
		if got == nil {
			t.Error("expected an error, but none was returned")
			return
		}

		// assert
		if !(got.Error() == expect.Error()) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		got := &[]*m.User{
			m.NewUser("Jenny", "Doe", "jenny_doe", "jenny@doe.com", "password1"),
			m.NewUser("Jane", "Doe", "jane_doe", "jane@doe.com", "password2"),
		}
		expect := &[]*m.User{
			{Id: 2, FirstName: "Jenny", LastName: "Doe", Username: "jenny_doe", Email: "jenny@doe.com"},
			{Id: 3, FirstName: "Jane", LastName: "Doe", Username: "jane_doe", Email: "jane@doe.com"},
		}

		// act
		err := app.HttpErrorFmts.User.AddAll(1, got)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})
}
