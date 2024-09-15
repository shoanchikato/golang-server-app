package author

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	v "app/pkg/validation"
	"app/test/setup"
	"reflect"
	"testing"
)

func addAuthor(t *testing.T, validation v.Validators) {
	author := m.NewAuthor("John", "Doe")

	err := validation.Author.Add(author)
	if err != nil {
		t.Error("error adding an author", err)
		return
	}
}

func Test_Author_Positive_Test(t *testing.T) {
	app := setup.Run()
	addAuthor(t, app.Validators)
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 1

		// act
		authors, err := app.HttpErrorFmts.Author.GetAll(1, 0, 50)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}
		got := len(*authors)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := m.Author{Id: 1, FirstName: "John", LastName: "Doe"}

		// act
		got, err := app.HttpErrorFmts.Author.GetOne(1, 1)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
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
		got := m.NewAuthor("John", "Doe")
		expect := m.Author{Id: 2, FirstName: "John", LastName: "Doe"}

		// act
		err := app.HttpErrorFmts.Author.Add(1, got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
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
		got := m.NewAuthor("John1", "Doe1")
		expect := m.Author{Id: 2, FirstName: "John1", LastName: "Doe1"}

		// act
		err := app.HttpErrorFmts.Author.Edit(1, 2, got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
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
		expect := e.NewErrRepoNotFound("author id", "2")

		// act
		err := app.HttpErrorFmts.Author.Remove(1, 2)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}
		_, got := app.HttpErrorFmts.Author.GetOne(1, 2)
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
		got := &[]*m.Author{
			m.NewAuthor("John", "Doe"),
			m.NewAuthor("James", "Doe"),
		}
		expect := &[]*m.Author{
			{Id: 2, FirstName: "John", LastName: "Doe", Books: nil},
			{Id: 3, FirstName: "James", LastName: "Doe", Books: nil},
		}

		// act
		err := app.HttpErrorFmts.Author.AddAll(1, got)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}

		// assert
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})
}
