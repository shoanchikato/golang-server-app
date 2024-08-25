package auth

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
) AuthValidator {
	repo := NewAuthRepo(db, rw, dbU)
	encrypt := NewAuthEncryption(repo, en)
	val := NewAuthValidator(encrypt)

	return val
}
