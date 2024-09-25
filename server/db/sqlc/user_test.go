package db

import (
	"context"
	"database/sql"
	"server/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createRandomUser(t *testing.T) (User, string) {
	pwd := utils.RandomString(12)
	arg := CreateUserParams{
		Username:       utils.RandomString(12),
		HashedPassword: pwd,
		FullName:       utils.RandomString(12),
		Email:          utils.RandomString(12),
	}

	arg.HashedPassword, _ = utils.HashPassword(arg.HashedPassword)

	user, err := testQueries.CreateUser(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, arg.Username, user.Username)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.Equal(t, arg.FullName, user.FullName)
	assert.Equal(t, arg.Email, user.Email)
	assert.NotZero(t, user.CreatedAt)

	_ = utils.CheckPassword(pwd, user.HashedPassword)

	return user, pwd
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1, pwd := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	assert.NoError(t, err)
	assert.NotEmpty(t, user2)

	assert.Equal(t, user1.Username, user2.Username)
	assert.Equal(t, user1.HashedPassword, user2.HashedPassword)
	assert.Equal(t, user1.FullName, user2.FullName)
	assert.Equal(t, user1.Email, user2.Email)
	assert.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	assert.Equal(t, user1.Role, user2.Role)
	_ = utils.CheckPassword(pwd, user2.HashedPassword)
}

func TestUpdateUser(t *testing.T) {
	user1, pwd := createRandomUser(t)

	arg := UpdateUserParams{
		HashedPassword:    sql.NullString{String: utils.RandomString(10), Valid: true},
		PasswordChangedAt: sql.NullTime{Time: time.Now(), Valid: true},
		FullName:          sql.NullString{String: utils.RandomString(10), Valid: true},
		Email:             sql.NullString{String: utils.RandomString(10), Valid: true},
		Username:          user1.Username,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, user2)

	assert.Equal(t, user1.Username, user2.Username)
	assert.Equal(t, arg.HashedPassword.String, user2.HashedPassword)
	assert.Equal(t, arg.FullName.String, user2.FullName)
	assert.Equal(t, arg.Email.String, user2.Email)
	assert.NotEqual(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	_ = utils.CheckPassword(pwd, user2.HashedPassword)
}
