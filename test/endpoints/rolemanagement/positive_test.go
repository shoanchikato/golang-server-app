package rolemanagement

import (
	"app/test/setup"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	m "app/pkg/model"
	v "app/pkg/validation"
)

func addRole(t *testing.T, validation v.Validators) {
	role := m.NewRole("default user")
	err := validation.Role.Add(role)
	if err != nil {
		t.Error(setup.UnexpectedErrorMsg, err)
		return
	}
}

func Test_Role_Management_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App
	validation := di.Validators
	addRole(t, validation)

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	t.Run("AddRoleToUser", func(t *testing.T) {
		// arrange
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodPost, "/role-management/2/1", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("GetRolesByUserId", func(t *testing.T) {
		// arrange
		got := &[]m.Role{}
		expect := m.Role{Id: 1, Name: "admin"}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/role-management/user/1", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
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

		if reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, *got)
			return
		}
	})

	t.Run("RemoveRoleFromUser", func(t *testing.T) {
		// arrange
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodPost, "/role-management/2/1", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

}
