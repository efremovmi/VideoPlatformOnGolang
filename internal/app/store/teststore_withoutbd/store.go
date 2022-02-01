package teststore_withoutbd

import (
	"github.com/genridarkbkru/registration/internal/app/model"
	"github.com/genridarkbkru/registration/internal/app/store"
)

// Структура БД.

// Store ...
type Store struct {
	userRepository *UserRepository
}

// Метод, который инициализирует БД.

// New ...
func New() *Store {
	return &Store{}
}

// Метод, не позволяющий использовать репозитории в обход БД.

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
