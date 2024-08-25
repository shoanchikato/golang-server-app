package repo

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	s "app/pkg/repo/stmt"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
)

type PostRepo interface {
	Add(post *m.Post) error
	AddAll(post *[]*m.Post) error
	Edit(id int, post *m.Post) error
	GetAll() (*[]m.Post, error)
	GetOne(id int) (*m.Post, error)
	Remove(id int) error
}

type postRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
}

func NewPostRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil) PostRepo {
	_, err := db.Exec(s.CREATE_POST_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating post table: %v: %s\n", err, s.CREATE_POST_TABLE_STMT)
	}

	return &postRepo{db, rw, dbU}
}

// Add
func (p *postRepo) Add(post *m.Post) error {
	id, err := p.dbU.Transaction(s.ADD_POST_STMT, post.Title, post.Body, post.UserID)
	if err != nil {
		return errors.Join(e.ErrRepoAdd, err)
	}

	post.ID = int(id)

	return nil
}

// AddAll
func (p *postRepo) AddAll(posts *[]*m.Post) error {
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
func (p *postRepo) Edit(id int, post *m.Post) error {
	_, err := p.dbU.Transaction(s.EDIT_POST_STMT, post.Title, post.Body, post.UserID, id)
	if err != nil {
		return errors.Join(e.ErrRepoEdit, err)
	}

	return nil
}

// GetAll
func (p *postRepo) GetAll() (*[]m.Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	post := m.Post{}
	posts := []m.Post{}

	rows, err := p.db.Query(s.GET_ALL_POST_STMT)
	if err != nil {
		return nil, errors.Join(e.ErrRepoGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
		if err != nil {
			return nil, errors.Join(e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
		}

		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Join(e.ErrRepoGetAll, e.ErrRepoExecutingStmt, err)
	}

	return &posts, nil
}

// GetOne
func (p *postRepo) GetOne(id int) (*m.Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	post := m.Post{}

	row := p.db.QueryRow(s.GET_ONE_POST_STMT, id)
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrRepoExecutingStmt, e.NewErrRepoNotFound(strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrRepoGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &post, nil
}

// Remove
func (p *postRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(s.REMOVE_POST_STMT, id)
	if err != nil {
		return errors.Join(e.ErrRepoRemove, err)
	}

	return nil
}
