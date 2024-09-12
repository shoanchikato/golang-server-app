package permissionmanagement

import (
	m "app/pkg/model"
	v "app/pkg/validation"
	"app/test/setup"
	"reflect"
	"testing"
)

type PermissionManagementHttpErrorFmt interface {
	AddPermissionToRole(adminId int, permissionId, roleId int) error
	AddPermissionsToRole(adminId int, permissionIds []int, roleId int) error
	GetPermissionsByRoleId(adminId int, roleId int) (*[]m.Permission, error)
	GetPermissionsByUserId(adminId int, userId int) (*[]m.Permission, error)
	RemovePermissionFromRole(adminId int, roleId, permissionId int) error
	RemovePermissionsFromRole(adminId int, roleId int, permissionIds []int) error
}

func addNewRole(t *testing.T, validation v.Validators) {
	role := m.NewRole("default user")
	err := validation.Role.Add(role)
	if err != nil {
		t.Error("error adding new role", err)
		return
	}
}

func TestPositivePermissionManagement(t *testing.T) {
	app := setup.Run()
	addNewRole(t, app.Validators)
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("AddPermissionToRole", func(t *testing.T) {
		// arrange
		permissionId := 1
		roleId := 2

		// act
		err := app.HttpErrorFmts.PermissionManagement.AddPermissionToRole(1, permissionId, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByRoleId(1, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if permissionId != (*permissions)[0].Id {
			t.Errorf("expected %v, got %v", permissionId, (*permissions)[1].Id)
			return
		}
	})

	t.Run("RemovePermissionFromRole", func(t *testing.T) {
		// arrange
		permissionId := 1
		roleId := 2

		// act
		err := app.HttpErrorFmts.PermissionManagement.RemovePermissionFromRole(1, roleId, permissionId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByRoleId(1, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if len(*permissions) != 0 {
			t.Errorf("expected %v, got %v", 0, len(*permissions))
			return
		}
	})

	t.Run("AddPermissionsToRole", func(t *testing.T) {
		// arrange
		permissionIds := []int{1, 2, 3}
		roleId := 2

		// act
		err := app.HttpErrorFmts.PermissionManagement.AddPermissionsToRole(1, permissionIds, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByRoleId(1, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		got := len(*permissions)
		expect := len(permissionIds)

		if expect != got {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetPermissionsByRoleId", func(t *testing.T) {
		// arrange
		expect := []int{1, 2, 3}
		roleId := 2

		// act
		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByRoleId(1, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		got := [3]int{}
		for i := range *permissions {
			got[i] = (*permissions)[i].Id
		}

		// assert
		if reflect.DeepEqual(got, expect) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetPermissionsByUserId", func(t *testing.T) {
		// arrange
		expect := 48
		userId := 1

		// act
		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByUserId(1, userId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		got := len(*permissions)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("RemovePermissionsFromRole", func(t *testing.T) {
		// arrange
		permissionIds := []int{1, 2, 3}
		roleId := 2

		// act
		err := app.HttpErrorFmts.PermissionManagement.RemovePermissionsFromRole(1, roleId, permissionIds)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		permissions, err := app.HttpErrorFmts.PermissionManagement.GetPermissionsByRoleId(1, roleId)
		if err != nil {
			t.Error("got unexpected error in test:", err)
			return
		}

		// assert
		if len(*permissions) != 0 {
			t.Errorf("expected %v, got %v", 0, len(*permissions))
			return
		}
	})

}
