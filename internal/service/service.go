package service

import (
	"songs-treasure/pkg/db"
)

type Service struct {
	db db.DB
}

func NewService(db db.DB) *Service {
	return &Service{
		db: db,
	}
}
