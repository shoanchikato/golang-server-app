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
}

func NewBookRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) BookRepo {

	return &bookRepo{db, rw, dbU}
}

// Add
func (p *bookRepo) Add(book *m.Book) error {
	id, err := p.dbU.Transaction(st.ADD_BOOK_STMT, book.Name, book.Year)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	book.Id = int(id)

	_, err = p.dbU.Transaction(st.ADD_AUTHOR_BOOK_RLTN_STMT, book.AuthorId, book.Id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return nil
}

// AddAll
func (p *bookRepo) AddAll(books *[]*m.Book) error {
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
func (p *bookRepo) Edit(id int, book *m.Book) error {
	idx, err := p.dbU.Transaction(st.EDIT_BOOK_STMT, book.Name, book.Year, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnEdit, err)
	}

	book.Id = int(idx)

	return nil
}

// GetAll
func (p *bookRepo) GetAll(lastId, limit int) (*[]m.Book, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	book := m.Book{}
	books := []m.Book{}

	rows, err := p.db.Query(st.GET_ALL_BOOK_STMT, lastId, limit)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Name, &book.Year)
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
	err := row.Scan(&book.Id, &book.Name, &book.Year)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrBookDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &book, nil
}

// Remove
func (p *bookRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(st.REMOVE_BOOK_STMT, id)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnRemove, err)
	}

	return nil
}
