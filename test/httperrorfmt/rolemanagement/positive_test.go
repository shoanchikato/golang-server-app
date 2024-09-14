package rolemanagement

import (
	"app/test/setup"
	"testing"

	e "app/pkg/errors"
	m "app/pkg/model"
	v "app/pkg/validation"
)

func addNewRole(t *testing.T, validation v.Validators) {
	role := m.NewRole("super user")
	err := validation.Role.Add(role)
	if err != nil {
		t.Error("error adding new role", err)
		return
	}

	user := m.NewUser("James", "Doe", "james_doe", "james@doe.com", "password2")
	err = validation.User.Add(user)
	if err != nil {
		t.Error("error adding new user:", err)
		return
	}
}

func TestPositiveRoleManagement(t *testing.T) {
	app := setup.Run()
	addNewRole(t, app.Validators)
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("AddRoleToUser", func(t *testing.T) {
		// arrange
		roleId := 2
		userId := 2
		expect := m.Role{Id: 2, Name: "super user"}

		// act
		err := app.HttpErrorFmts.RoleManagement.AddRoleToUser(1, roleId, userId)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		got, err := app.HttpErrorFmts.RoleManagement.GetRoleByUserId(1, userId)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if expect != *got {
			t.Errorf("expected %v, got %v", expect, *got)
			return
		}
	})

	t.Run("GetRoleByUserId", func(t *testing.T) {
		// arrange
		expect := m.Role{Id: 2, Name: "super user"}
		userId := 2

		// act
		got, err := app.HttpErrorFmts.RoleManagement.GetRoleByUserId(1, userId)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if *got != expect {
			t.Errorf("expected %v, got %v", expect, *got)
			return
		}
	})

	t.Run("RemoveRoleFromUser", func(t *testing.T) {
		// arrange
		roleId := 2
		userId := 2
		expect := e.NewErrRepoNotFound("role id", "2")

		// act
		err := app.HttpErrorFmts.RoleManagement.RemoveRoleFromUser(1, roleId, userId)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		_, got := app.HttpErrorFmts.RoleManagement.GetRoleByUserId(1, userId)

		// assert
		if expect.Error() != got.Error() {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})
}
