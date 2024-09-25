package db

import (
	"context"
	"server/token"
	"server/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	user, _ := createRandomUser(t)
	symmetricKey := utils.RandomString(32)
	maker, err := token.NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)

	sessionID := uuid.New()
	username := user.Username
	role := user.Role
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, role, duration)
	assert.NoError(t, err)

	createSessionParams := CreateSessionParams{
		ID:           sessionID,
		Username:     username,
		RefreshToken: token,
		UserAgent:    "Mozilla/5.0",
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    payload.ExpiredAt,
	}

	session, err := testQueries.CreateSession(context.Background(), createSessionParams)

	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, sessionID, session.ID)
	assert.Equal(t, username, session.Username)
	assert.Equal(t, token, session.RefreshToken)
}

func TestVerifySessionToken(t *testing.T) {
	symmetricKey := utils.RandomString(32)
	user, _ := createRandomUser(t)
	maker, err := token.NewPasetoMaker(symmetricKey)
	assert.NoError(t, err)

	username := user.Username
	role := user.Role
	duration := time.Minute
	token, payload, err := maker.CreateToken(username, role, duration)
	assert.NoError(t, err)

	verifiedPayload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, verifiedPayload)
	assert.Equal(t, payload.Username, verifiedPayload.Username)
	assert.Equal(t, payload.Role, verifiedPayload.Role)
}
