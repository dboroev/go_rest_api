package store_test

import (
    "github.com/stretchr/testify/assert"
    "go/http-rest-api/internal/app/model"
    "go/http-rest-api/internal/app/store"
    "testing"
)

func TestUserRepository_Create(t *testing.T) {
    s, teardown := store.TestStore(t, databaseUrl)
    defer teardown("users")

    u, err := s.User().Create(model.TestUser(t))

    assert.NoError(t, err)
    assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
    s, teardown := store.TestStore(t, databaseUrl)
    defer teardown("users")

    email := "user@example.com"
    _, err := s.User().FindByEmail(email)

    assert.Error(t, err)

    u := model.TestUser(t)
    s.User().Create(u)

    findUser, err := s.User().FindByEmail(u.Email)

    assert.NoError(t, err)
    assert.NotNil(t, findUser)
}
