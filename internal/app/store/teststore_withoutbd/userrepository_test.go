package teststore_withoutbd_test

import (
	"github.com/genridarkbkru/registration/internal/app/model"
	"github.com/genridarkbkru/registration/internal/app/store"
	"github.com/genridarkbkru/registration/internal/app/store/teststore_withoutbd"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Тестирование метода Create у UserRepository.

func TestUserRepository_Create(t *testing.T) {
	s := teststore_withoutbd.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

// Тестирование метода FindByEmail у UserRepository.

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore_withoutbd.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email) // Попытка найти email.
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	// Создание поле в БД.
	s.User().Create(u)

	u, err = s.User().FindByEmail(email) // Попытка найти email.

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
