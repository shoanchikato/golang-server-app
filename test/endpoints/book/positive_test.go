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
	v "app/pkg/validation"
)

func addAuthor(t *testing.T, validation v.Validators) {
	author := m.NewAuthor("John", "Doe")
	err := validation.Author.Add(author)
	if err != nil {
		t.Error(setup.UnexpectedErrorMsg, err)
		return
	}
}

func Test_Book_Endpoint__Positive_Test(t *testing.T) {
	di := setup.Run()
	app := di.App
	validation := di.Validators
	addAuthor(t, validation)

	tokens, err := setup.GetAuthTokens(app)
	if err != nil {
		t.Error(setup.UnexpectedErrorMsg, err)
		return
	}

	t.Run("Add", func(t *testing.T) {
		// arrange
		value := `{"name":"book one", "year":2012, "author_id":1}`
		reader := strings.NewReader(value)
		expect := &m.Book{Id: 1, Name: "book one", Year: 2012, AuthorId: 1}
		got := &m.Book{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/books", reader)
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

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("AddAll", func(t *testing.T) {
		// arrange
		value := `
			[
				{"name":"book two", "year":2012, "author_id":1},
				{"name":"book three", "year":2013, "author_id":1}
			]
		`
		reader := strings.NewReader(value)
		expect := &[]m.Book{
			{Id: 2, Name: "book two", Year: 2012, AuthorId: 1},
			{Id: 3, Name: "book three", Year: 2013, AuthorId: 1},
		}
		got := &[]m.Book{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPost, "/books/all", reader)
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

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Remove", func(t *testing.T) {
		// arrange
		expectStatus := http.StatusNoContent

		// act
		req := httptest.NewRequest(http.MethodDelete, "/books/2", nil)
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

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 2
		got := &[]m.Book{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
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

		if expect != len(*got) {
			t.Errorf("expected %v, but got %v", expect, len(*got))
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := &m.Book{Id: 1, Name: "book one", Year: 2012, AuthorId: 1}
		got := &m.Book{}
		expectStatus := http.StatusOK

		// act
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
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

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})

	t.Run("Edit", func(t *testing.T) {
		// arrange
		value := `{"name":"book three", "year":2023, "author_id":1}`
		reader := strings.NewReader(value)
		expect := &m.Book{Id: 1, Name: "book three", Year: 2023, AuthorId: 1}
		got := &m.Book{}
		expectStatus := http.StatusCreated

		// act
		req := httptest.NewRequest(http.MethodPut, "/books/1", reader)
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

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, but got %v", expect, got)
			return
		}
	})
}
