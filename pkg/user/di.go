package user

import (
	r "app/pkg/repo"
	s "app/pkg/service"
	"database/sql"
	"sync"
)

func Di(
	db *sql.DB,
	rw *sync.RWMutex,
	dbU r.DBUtil,
	en s.EncryptionService,
) UserValidator {
	repo := NewUserRepo(db, rw, dbU)
	encrypt := NewUserEncryption(repo, en)
	val := NewUserValidator(encrypt)

	return val
}
