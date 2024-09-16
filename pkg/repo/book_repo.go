package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	st "app/pkg/stmt"
	"database/sql"
	"errors"
	"strconv"
	"sync"
)

type BookRepo interface {
	Add(book *m.Book) error
	AddAll(book *[]*m.Book) error
	Edit(id int, book *m.Book) error
	GetAll(lastId, limit int) (*[]m.Book, error)
	GetOne(id int) (*m.Book, error)
	Remove(id int) error
}

type bookRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
	ar  AuthorRepo
}

func NewBookRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil, ar AuthorRepo) BookRepo {

	return &bookRepo{db, rw, dbU, ar}
}

// Add
func (p *bookRepo) Add(book *m.Book) error {
	_, err := p.ar.GetOne(book.AuthorId)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	id, err := p.dbU.Transaction(st.ADD_BOOK_STMT, book.Name, book.Year, book.AuthorId)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	book.Id = int(id)

	return nil
}

// AddAll
func (p *bookRepo) AddAll(books *[]*m.Book) error {
	for _, book := range *books {
		err := p.Add(book)
		if err != nil {
			return errors.Join(e.ErrBookDomain, e.ErrOnAddAll, err)
		}
	}

	return nil
}

// Edit
func (p *bookRepo) Edit(id int, book *m.Book) error {
	_, err := p.dbU.Transaction(st.EDIT_BOOK_STMT, book.Name, book.Year, id, book.AuthorId, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnEdit, err)
	}

	book.Id = id

	return nil
}

// GetAll
func (p *bookRepo) GetAll(lastId, limit int) (*[]m.Book, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	p.dbU.CheckLimit(&limit)

	book := m.Book{}
	books := []m.Book{}

	rows, err := p.db.Query(st.GET_ALL_BOOK_STMT, lastId, limit)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Name, &book.Year, &book.AuthorId)
		if err != nil {
			return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &books, nil
}

// GetOne
func (p *bookRepo) GetOne(id int) (*m.Book, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	book := m.Book{}

	row := p.db.QueryRow(st.GET_ONE_BOOK_STMT, id)
	err := row.Scan(&book.Id, &book.Name, &book.Year, &book.AuthorId)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound("book id", strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &book, nil
}

// Remove
func (p *bookRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(st.REMOVE_BOOK_STMT, id, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnRemove, err)
	}

	return nil
}
