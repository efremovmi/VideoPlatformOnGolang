package sqlstore

import (
	"database/sql"
	"github.com/genridarkbkru/registration/internal/app/store"
	_ "github.com/lib/pq" // ...
)

// Структура БД.

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// Метод, который инициализирует БД.

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Метод, не позволяющий использовать репозитории в обход БД.

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
