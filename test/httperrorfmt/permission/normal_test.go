package permission

import (
	"app/pkg/di"
	e "app/pkg/errors"
	m "app/pkg/model"
	"app/test/setup"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestPermission(t *testing.T) {
	app := setup.Run()
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 48

		// act
		permissions, err := app.HttpErrorFmts.Permission.GetAll(1, 0, 50)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}
		got := len(*permissions)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		permissions, _, _, _, _ := di.Data()
		permission := permissions[0]
		permission.Id = 1
		expect := permission

		// act
		got, err := app.HttpErrorFmts.Permission.GetOne(1, 1)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if *got != *expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})

	t.Run("Add", func(t *testing.T) {
		// arrange
		got := m.NewPermission("name", "entity", "operation")
		expect := m.Permission{Id: 49, Name: "name", Entity: "entity", Operation: "operation"}

		// act
		err := app.HttpErrorFmts.Permission.Add(1, got)
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
		got := m.NewPermission("name1", "entity1", "operation1")
		expect := m.Permission{Id: 49, Name: "name1", Entity: "entity1", Operation: "operation1"}

		// act
		err := app.HttpErrorFmts.Permission.Edit(1, 49, got)
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
		expect := e.NewErrRepoNotFound("permission id", "49")

		// act
		err := app.HttpErrorFmts.Permission.Remove(1, 49)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}
		_, got := app.HttpErrorFmts.Permission.GetOne(1, 49)
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
		got := &[]*m.Permission{
			m.NewPermission("name", "entity", "operation"),
			m.NewPermission("name", "entity", "operation"),
		}
		expect := &[]*m.Permission{
			{Id: 49, Name: "name", Entity: "entity", Operation: "operation"},
			{Id: 50, Name: "name", Entity: "entity", Operation: "operation"},
		}

		// act
		err := app.HttpErrorFmts.Permission.AddAll(1, got)
		if err != nil {
			t.Error("got unexpected error in test", err)
		}

		// assert
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetByEntity", func(t *testing.T) {
		// arrange
		expect := &[]m.Permission{
			{Id: 1, Name: "auth: login", Entity: "auth", Operation: "login"},
			{Id: 2, Name: "auth: reset password", Entity: "auth", Operation: "reset password"},
		}

		// act
		got, err := app.HttpErrorFmts.Permission.GetByEntity(1, "auth")
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
