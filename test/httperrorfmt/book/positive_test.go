package book

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	v "app/pkg/validation"
	"app/test/setup"
	"fmt"
	"reflect"
	"testing"
)

func addAuthor(t *testing.T, validation v.Validators) {
	author := m.NewAuthor("John", "Doe")
	book := m.NewBook("book one", 2012, 1)

	err := validation.Author.Add(author)
	if err != nil {
		t.Error("error adding author", err)
		return
	}

	err = validation.Book.Add(book)
	if err != nil {
		fmt.Println("error adding book")
		return
	}
}

func Test_Book_Positive_Test(t *testing.T) {
	app := setup.Run()
	addAuthor(t, app.Validators)
	t.Cleanup(func() {
		setup.CleanUp(app)
	})

	t.Run("GetAll", func(t *testing.T) {
		// arrange
		expect := 1

		// act
		books, err := app.HttpErrorFmts.Book.GetAll(1, 0, 50)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}
		got := len(*books)

		// assert
		if got != expect {
			t.Errorf("expected %v, got %v", expect, got)
			return
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		// arrange
		expect := m.Book{Id: 1, Name: "book one", Year: 2012, AuthorId: 1}

		// act
		got, err := app.HttpErrorFmts.Book.GetOne(1, 1)
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
		got := m.NewBook("book two", 2013, 1)
		expect := m.Book{Id: 2, Name: "book two", Year: 2013, AuthorId: 1}

		// act
		err := app.HttpErrorFmts.Book.Add(1, got)
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
		got := m.NewBook("book 2", 2014, 1)
		expect := m.Book{Id: 2, Name: "book 2", Year: 2014, AuthorId: 1}

		// act
		err := app.HttpErrorFmts.Book.Edit(1, 2, got)
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
		expect := e.NewErrRepoNotFound("book id", "2")

		// act
		err := app.HttpErrorFmts.Book.Remove(1, 2)
		if err != nil {
			t.Error(setup.UnexpectedErrorMsg, err)
			return
		}
		_, got := app.HttpErrorFmts.Book.GetOne(1, 2)
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
		got := &[]*m.Book{
			m.NewBook("book two", 2012, 1),
			m.NewBook("book three", 2013, 1),
		}
		expect := &[]*m.Book{
			{Id: 2, Name: "book two", Year: 2012, AuthorId: 1},
			{Id: 3, Name: "book three", Year: 2013, AuthorId: 1},
		}

		// act
		err := app.HttpErrorFmts.Book.AddAll(1, got)
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
