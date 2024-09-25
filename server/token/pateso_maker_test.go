package token

import (
	"server/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPasetoMaker(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)
	assert.NotNil(t, maker)

	invalidKey := utils.RandomString(10)
	maker, err = NewPasetoMaker(invalidKey)
	assert.Error(t, err)
	assert.Nil(t, maker)
	assert.EqualError(t, err, "invalid key size: must be exactly 32 characters")
}

func TestCreateToken(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)

	username := utils.RandomString(10)
	role := "user"
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, role, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotNil(t, payload)
	assert.Equal(t, username, payload.Username)
	assert.Equal(t, role, payload.Role)
	assert.WithinDuration(t, time.Now().Add(duration), payload.ExpiredAt, time.Second)
}

func TestVerifyToken(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)

	username := utils.RandomString(10)
	role := "user"
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, role, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify the token
	verifiedPayload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, verifiedPayload)
	assert.Equal(t, payload.Username, verifiedPayload.Username)
	assert.Equal(t, payload.Role, verifiedPayload.Role)
	assert.WithinDuration(t, payload.ExpiredAt, verifiedPayload.ExpiredAt, time.Second)
}

func TestVerifyExpiredToken(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)

	username := utils.RandomString(10)
	role := "user"
	duration := -time.Minute

	token, _, err := maker.CreateToken(username, role, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	verifiedPayload, err := maker.VerifyToken(token)
	assert.Error(t, err)
	assert.Nil(t, verifiedPayload)
	assert.EqualError(t, err, ErrExpiredToken.Error())
}
