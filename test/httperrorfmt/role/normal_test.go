package role

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	"app/test/setup"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestRole(t *testing.T) {
	app := setup.Run()
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 1

		// act
		roles, err := app.HttpErrorFmts.Role.GetAll(1, 0, 50)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}
		got := len(*roles)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := m.Role{Id: 1, Name: "admin"}

		// act
		got, err := app.HttpErrorFmts.Role.GetOne(1, 1)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("Add", func(t *testing.T) {
		// arrange
		got := m.NewRole("name")
		expect := m.Role{Id: 2, Name: "name"}

		// act
		err := app.HttpErrorFmts.Role.Add(1, got)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		got := m.NewRole("name1")
		expect := m.Role{Id: 2, Name: "name1"}

		// act
		err := app.HttpErrorFmts.Role.Edit(1, 49, got)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expect := e.NewErrRepoNotFound("role id", "2")

		// act
		err := app.HttpErrorFmts.Role.Remove(1, 2)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}
		_, got := app.HttpErrorFmts.Role.GetOne(1, 2)
		if got == nil {
			t.Error("expected an error, but none was returned")
			return
		}

		// assert
		if !(got.Error() == expect.Error()) {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		got := &[]*m.Role{
			m.NewRole("name"),
			m.NewRole("name"),
		}
		expect := &[]*m.Role{
			{Id: 2, Name: "name"},
			{Id: 3, Name: "name"},
		}

		// act
		err := app.HttpErrorFmts.Role.AddAll(1, got)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})
}
