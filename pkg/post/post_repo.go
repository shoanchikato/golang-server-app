package post

import (
	e "app/pkg/errors"
	r "app/pkg/repo"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type PostRepo interface {
	Add(post *Post) error
	AddAll(post *[]*Post) error
	Edit(id int, post *Post) error
	GetAll() (*[]Post, error)
	GetOne(id int) (*Post, error)
	Remove(id int) error
}

type postRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU r.DBUtil
}

func NewPostRepo(db *sql.DB, rw *sync.RWMutex, dbU r.DBUtil) PostRepo {
	_, err := db.Exec(CREATE_POST_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating post table: %v: %s\n", err, CREATE_POST_TABLE_STMT)
	}

	return &postRepo{db, rw, dbU}
}

// Add
func (p *postRepo) Add(post *Post) error {
	id, err := p.dbU.Transaction(ADD_POST_STMT, post.Title, post.Body, post.UserID)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	post.ID = int(id)

	return nil
}

// AddAll
func (p *postRepo) AddAll(posts *[]*Post) error {
	newPosts := *posts
	for i := 0; i < len(newPosts); i++ {
		post := newPosts[i]
		err := p.Add(post)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit
func (p *postRepo) Edit(id int, post *Post) error {
	idx, err := p.dbU.Transaction(EDIT_POST_STMT, post.Title, post.Body, post.UserID, id)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnEdit, err)
	}

	post.ID = int(idx)

	return nil
}

// GetAll
func (p *postRepo) GetAll() (*[]Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	post := Post{}
	posts := []Post{}

	rows, err := p.db.Query(GET_ALL_POST_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
		if err != nil {
			return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
		}

		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &posts, nil
}

// GetOne
func (p *postRepo) GetOne(id int) (*Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	post := Post{}

	row := p.db.QueryRow(GET_ONE_POST_STMT, id)
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrPostDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &post, nil
}

// Remove
func (p *postRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(REMOVE_POST_STMT, id)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnRemove, err)
	}

	return nil
}
