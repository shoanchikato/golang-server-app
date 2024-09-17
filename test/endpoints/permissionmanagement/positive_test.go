package permissionmanagement

import (
	"app/test/setup"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	m "app/pkg/model"
	v "app/pkg/validation"
)

// permissionManagement := app.Group("/permission-management")
// 	permissionManagement.Use(middleware.JWTParser)
// 	permissionManagement.Get("/role/:roleId", handler.GetPermissionsByRoleId)
// 	permissionManagement.Get("/user/:userId", handler.GetPermissionsByUserId)
// 	permissionManagement.Post("role/:roleId", handler.AddPermissionsToRole)
// 	permissionManagement.Post("/:permissionId/:roleId", handler.AddPermissionToRole)
// 	permissionManagement.Delete("role/:roleId", handler.RemovePermissionsFromRole)
// 	permissionManagement.Delete(":permissionId/:roleId", handler.RemovePermissionFromRole)
// }

func addRole(t *testing.T, validation v.Validators) {
	role := m.NewRole("default user")
	err := validation.Role.Add(role)
	if err != nil {
		t.Errorf(setup.UnexpectedErrorMsg)
		return
	}
}

func Test_Permission_Management_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App
	validation := di.Validators
	addRole(t, validation)

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	t.Run("AddPermissionsToRole", func(t *testing.T) {
		// arrange
		value := `[1, 2, 3]`
		reader := strings.NewReader(value)
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodPost, "/permission-management/role/2", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("RemovePermissionsFromRole", func(t *testing.T) {
		// arrange
		value := `[1, 2, 3]`
		reader := strings.NewReader(value)
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodDelete, "/permission-management/role/2", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("AddPermissionToRole", func(t *testing.T) {
		// arrange
		roleId := 2
		permissionId := 4
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/permission-management/%d/%d", permissionId, roleId), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("RemovePermissionFromRole", func(t *testing.T) {
		// arrange
		roleId := 2
		permissionId := 4
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/permission-management/%d/%d", permissionId, roleId), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("GetPermissionsByRoleId", func(t *testing.T) {
		// arrange
		got := &[]m.Permission{}
		expect := 0
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/permission-management/role/2", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if expect != len(*got) {
			t.Errorf("expected %v, but got %v", expect, len(*got))
			return
		}
	})

	t.Run("GetPermissionsByUserId", func(t *testing.T) {
		// arrange
		got := &[]m.Permission{}
		expect := 48
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/permission-management/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if expect != len(*got) {
			t.Errorf("expected %v, but go %v", expect, len(*got))
			return
		}
	})

}
