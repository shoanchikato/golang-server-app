package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	st "app/pkg/stmt"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type AuthorRepo interface {
	Add(author *m.Author) error
	AddAll(author *[]*m.Author) error
	Edit(id int, author *m.Author) error
	GetAll() (*[]m.Author, error)
	GetOne(id int) (*m.Author, error)
	GetMore(id int) (*m.Author, error)
	Remove(id int) error
}

type authorRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewAuthorRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) AuthorRepo {
	_, err := db.Exec(st.CREATE_AUTHOR_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating author table: %v: %s\n", err, st.CREATE_AUTHOR_TABLE_STMT)
	}

	return &authorRepo{db, rw, dbU}
}

// Add
func (p *authorRepo) Add(author *m.Author) error {
	id, err := p.dbU.Transaction(st.ADD_AUTHOR_STMT, author.FirstName, author.LastName)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnAdd, err)
	}

	author.Id = int(id)

	return nil
}

// AddAll
func (p *authorRepo) AddAll(authors *[]*m.Author) error {
	newAuthors := *authors
	for i := 0; i < len(newAuthors); i++ {
		author := newAuthors[i]
		err := p.Add(author)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *authorRepo) Edit(id int, author *m.Author) error {
	idx, err := p.dbU.Transaction(st.EDIT_AUTHOR_STMT, author.FirstName, author.LastName, id)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnEdit, err)
	}

	author.Id = int(idx)

	return nil
}

// GetAll
func (p *authorRepo) GetAll() (*[]m.Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author := m.Author{}
	authors := []m.Author{}

	rows, err := p.db.Query(st.GET_ALL_AUTHOR_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&author.Id, &author.FirstName, &author.LastName)
		if err != nil {
			return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		authors = append(authors, author)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &authors, nil
}

// GetOne
func (p *authorRepo) GetOne(id int) (*m.Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author := m.Author{}

	row := p.db.QueryRow(st.GET_ONE_AUTHOR_STMT, id)
	err := row.Scan(&author.Id, &author.FirstName, &author.LastName)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &author, nil
}

func (p *authorRepo) GetMore(id int) (*m.Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author, err := p.GetOne(id)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, err)
	}

	book := m.Book{}
	books := []m.Book{}

	rows, err := p.db.Query(st.GET_BOOKS_BY_AUTHOR_Id_STMT, id)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Name, &book.Year, &book.AuthorId)
		if err != nil {
			return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, e.ErrRepoExecutingStmt, err)
		}

		books = append(books, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, e.ErrRepoExecutingStmt, err)
	}

	author.Books = &books

	return author, nil
}

// Remove
func (p *authorRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(st.REMOVE_AUTHOR_STMT, id)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnRemove, err)
	}

	return nil
}
