package store

import "github.com/genridarkbkru/registration/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
