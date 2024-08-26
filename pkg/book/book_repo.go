package book

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"strconv"
	"sync"
)

type BookRepo interface {
	Add(book *Book) error
	AddAll(book *[]*Book) error
	Edit(id int, book *Book) error
	GetAll() (*[]Book, error)
	GetOne(id int) (*Book, error)
	Remove(id int) error
}

type bookRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewBookRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) BookRepo {

	return &bookRepo{db, rw, dbU}
}

// Add
func (p *bookRepo) Add(book *Book) error {
	id, err := p.dbU.Transaction(ADD_BOOK_STMT, book.Name, book.Year)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrRepoAdd, err)
	}

	book.ID = int(id)

	_, err = p.dbU.Transaction(ADD_AUTHOR_BOOK_RLTN_STMT, book.AuthorID, book.ID)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrRepoAdd, err)
	}

	return nil
}

// AddAll
func (p *bookRepo) AddAll(books *[]*Book) error {
	newBooks := *books
	for i := 0; i < len(newBooks); i++ {
		book := newBooks[i]
		err := p.Add(book)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *bookRepo) Edit(id int, book *Book) error {
	_, err := p.dbU.Transaction(EDIT_BOOK_STMT, book.Name, book.Year, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrRepoEdit, err)
	}

	return nil
}

// GetAll
func (p *bookRepo) GetAll() (*[]Book, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	book := Book{}
	books := []Book{}

	rows, err := p.db.Query(GET_ALL_BOOK_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Name, &book.Year)
		if err != nil {
			return nil, errors.Join(e.ErrBookDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &books, nil
}

// GetOne
func (p *bookRepo) GetOne(id int) (*Book, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	book := Book{}

	row := p.db.QueryRow(GET_ONE_BOOK_STMT, id)
	err := row.Scan(&book.ID, &book.Name, &book.Year)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &book, nil
}

// Remove
func (p *bookRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(REMOVE_BOOK_STMT, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrRepoRemove, err)
	}

	return nil
}
