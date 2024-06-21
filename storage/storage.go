package storage

import (
	"github.com/SaidovZohid/medium-clone-simple/storage/postgres"
	"github.com/SaidovZohid/medium-clone-simple/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStorage(psqlConn *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUserStorage(psqlConn),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
