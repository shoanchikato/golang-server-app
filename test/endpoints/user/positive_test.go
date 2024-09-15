package endpoints

import (
	m "app/pkg/model"
	"app/test/setup"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_User_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error("unexpected error", err)
		return
	}

	t.Run("Add", func(t *testing.T) {
		// arrange
		value := `
			{
				"first_name":"James", 
				"last_name":"Doe", 
				"username":"james_doe", 
				"email":"james@doe.com", 
				"password":"password1"
			}
		`
		reader := strings.NewReader(value)
		expect := &m.User{
			Id: 2, FirstName: "James", 
			LastName: "Doe", 
			Username: "james_doe", 
			Email: "james@doe.com",
		}
		got := &m.User{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/users", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		value := `
			[
				{
					"first_name":"Joe", 
					"last_name":"Doe", 
					"username":"joe_doe", 
					"email":"joe@doe.com", 
					"password":"password1"
				},
				{
					"first_name":"Jamie", 
					"last_name":"Doe", 
					"username":"jamie_doe", 
					"email":"jamie@doe.com", 
					"password":"password1"
				}
			]
		`
		reader := strings.NewReader(value)
		expect := &[]m.User{
			{
				Id: 3, FirstName: "Joe", 
				LastName: "Doe", 
				Username: "joe_doe", 
				Email: "joe@doe.com",
			},
			{
				Id: 4, FirstName: "Jamie", 
				LastName: "Doe", 
				Username: "jamie_doe", 
				Email: "jamie@doe.com",
			},
		}
		got := &[]m.User{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/users/all", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodDelete, "/users/3", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 3
		got := &[]m.User{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if expect != len(*got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := &m.User{
			Id: 2, FirstName: "James", 
			LastName: "Doe", 
			Username: "james_doe", 
			Email: "james@doe.com",
		}
		got := &m.User{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		value := `
			{
				"first_name":"James1", 
				"last_name":"Doe1", 
				"username":"james_doe1", 
				"email":"james@doe.com1"
			}
		`
		reader := strings.NewReader(value)
		expect := &m.User{
			Id: 2, FirstName: "James1", 
			LastName: "Doe1", 
			Username: "james_doe1", 
			Email: "james@doe.com1",
		}
		got := &m.User{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPut, "/users/2", reader)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access)

		resp, err := app.Test(req)
		if err != nil {
			t.Error("unexpected error", err)
			return
		}

		// assert
		if resp.StatusCode != expectStatus {
			errResp := setup.GetErrorResponse(t, resp.Body)
			t.Errorf("expected %v, but got %v, response %v", expectStatus, resp.StatusCode, errResp)
			return
		}

		err = json.NewDecoder(resp.Body).Decode(got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})
}
