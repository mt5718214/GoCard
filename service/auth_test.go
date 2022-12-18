package service

import (
	"context"
	"gocard/db"
	sqlc "gocard/db/sqlc"
	enum "gocard/enum"
	"gocard/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRegisterHandler(t *testing.T) {
	username := util.RandomString(6)
	password := util.RandomPassword()
	email := util.RandomEmail()

	result, err := RegisterHandler(username, password, email)
	require.NoError(t, err)
	require.Equal(t, "Create user success", result)

	user, err := db.Queries.GetUserByEmail(context.Background(), email)
	require.NoError(t, err)
	require.Equal(t, username, user.Name)
	require.Equal(t, email, user.Email)

	require.True(t, util.CheckPasswordHash(user.Password, password))
	require.NotZero(t, user.ID)
}

func TestAuthHandler(t *testing.T) {
	password := util.RandomPassword()
	arg := sqlc.PostUserParams{
		Name:          util.RandomUser(),
		Email:         util.RandomEmail(),
		Password:      util.HashPassword(password),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}
	user, err := db.Queries.PostUser(context.Background(), arg)
	require.NoError(t, err)

	token, err := AuthHandler(user.Email, password)
	require.NoError(t, err)

	claim, err := parseToken(token)
	require.NoError(t, err)
	require.Equal(t, claim["sub"], "token")
	require.Equal(t, claim["jti"], uuid.UUID.String(user.ID))
	require.True(t, claim.VerifyAudience(user.Name, true))
}
