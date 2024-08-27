package author

import (
	b "app/pkg/book"
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type AuthorRepo interface {
	Add(author *Author) error
	AddAll(author *[]*Author) error
	Edit(id int, author *Author) error
	GetAll() (*[]Author, error)
	GetOne(id int) (*Author, error)
	GetMore(id int) (*Author, error)
	Remove(id int) error
}

type authorRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewAuthorRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) AuthorRepo {
	_, err := db.Exec(CREATE_AUTHOR_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating author table: %v: %s\n", err, CREATE_AUTHOR_TABLE_STMT)
	}

	return &authorRepo{db, rw, dbU}
}

// Add
func (p *authorRepo) Add(author *Author) error {
	id, err := p.dbU.Transaction(ADD_AUTHOR_STMT, author.FirstName, author.LastName)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnAdd, err)
	}

	author.ID = int(id)

	return nil
}

// AddAll
func (p *authorRepo) AddAll(authors *[]*Author) error {
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
func (p *authorRepo) Edit(id int, author *Author) error {
	idx, err := p.dbU.Transaction(EDIT_AUTHOR_STMT, author.FirstName, author.LastName, id)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnEdit, err)
	}

	author.ID = int(idx)

	return nil
}

// GetAll
func (p *authorRepo) GetAll() (*[]Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author := Author{}
	authors := []Author{}

	rows, err := p.db.Query(GET_ALL_AUTHOR_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.FirstName, &author.LastName)
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
func (p *authorRepo) GetOne(id int) (*Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author := Author{}

	row := p.db.QueryRow(GET_ONE_AUTHOR_STMT, id)
	err := row.Scan(&author.ID, &author.FirstName, &author.LastName)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &author, nil
}

func (p *authorRepo) GetMore(id int) (*Author, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	author, err := p.GetOne(id)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, err)
	}

	book := b.Book{}
	books := []b.Book{}

	rows, err := p.db.Query(GET_BOOKS_BY_AUTHOR_ID_STMT, id)
	if err != nil {
		return nil, errors.Join(e.ErrAuthDomain, e.ErrOnGetMore, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Name, &book.Year, &book.AuthorID)
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
	_, err := p.dbU.Transaction(REMOVE_AUTHOR_STMT, id)
	if err != nil {
		return errors.Join(e.ErrAuthDomain, e.ErrOnRemove, err)
	}

	return nil
}
