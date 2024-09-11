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

type PostRepo interface {
	Add(post *m.Post) error
	AddAll(post *[]*m.Post) error
	Edit(id int, post *m.Post) error
	GetAll(lastId, limit int) (*[]m.Post, error)
	GetOne(id int) (*m.Post, error)
	Remove(id int) error
}

type postRepo struct {
	db  *sql.DB
	rw  *sync.RWMutex
	dbU DBUtil
	ur  UserRepo
}

func NewPostRepo(db *sql.DB, rw *sync.RWMutex, dbU DBUtil, ur UserRepo) PostRepo {
	_, err := db.Exec(st.CREATE_POST_TABLE_STMT)
	if err != nil {
		log.Fatalf("error creating post table: %v: %s\n", err, st.CREATE_POST_TABLE_STMT)
	}

	return &postRepo{db, rw, dbU, ur}
}

// Add
func (p *postRepo) Add(post *m.Post) error {
	_, err := p.ur.GetOne(post.UserId)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	id, err := p.dbU.Transaction(st.ADD_POST_STMT, post.Title, post.Body, post.UserId)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnAdd, err)
	}

	post.Id = int(id)

	return nil
}

// AddAll
func (p *postRepo) AddAll(posts *[]*m.Post) error {
	for _, post := range *posts {
		err := p.Add(post)
		if err != nil {
			return errors.Join(e.ErrPostDomain, e.ErrOnAddAll, err)
		}
	}

	return nil
}

// Edit
func (p *postRepo) Edit(id int, post *m.Post) error {
	idx, err := p.dbU.Transaction(st.EDIT_POST_STMT, post.Title, post.Body, post.UserId, id)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnEdit, err)
	}

	post.Id = int(idx)

	return nil
}

// GetAll
func (p *postRepo) GetAll(lastId, limit int) (*[]m.Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	p.dbU.CheckLimit(&limit)

	post := m.Post{}
	posts := []m.Post{}

	rows, err := p.db.Query(st.GET_ALL_POST_STMT, lastId, limit)
	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetAll, e.ErrRepoPreparingStmt, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.UserId)
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
func (p *postRepo) GetOne(id int) (*m.Post, error) {
	p.rw.RLock()
	defer p.rw.RUnlock()

	post := m.Post{}

	row := p.db.QueryRow(st.GET_ONE_POST_STMT, id)
	err := row.Scan(&post.Id, &post.Title, &post.Body, &post.UserId)
	if err == sql.ErrNoRows {
		return nil, errors.Join(e.ErrPostDomain, e.ErrRepoExecutingStmt, e.NewErrRepoNotFound("post id", strconv.Itoa(id)))
	}

	if err != nil {
		return nil, errors.Join(e.ErrPostDomain, e.ErrOnGetOne, e.ErrRepoExecutingStmt, err)
	}

	return &post, nil
}

// Remove
func (p *postRepo) Remove(id int) error {
	_, err := p.dbU.Transaction(st.REMOVE_POST_STMT, id)
	if err != nil {
		return errors.Join(e.ErrPostDomain, e.ErrOnRemove, err)
	}

	return nil
}
