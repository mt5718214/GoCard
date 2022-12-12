package db

import (
	"context"
	"gocard/enum"
	"gocard/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	password := util.RandomPassword()
	arg := PostUserParams{
		Name:          util.RandomUser(),
		Email:         util.RandomEmail(),
		Password:      util.HashPassword(password),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}
	user, err := testQueries.PostUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.LastUpdatedBy, user.LastUpdatedBy)
	require.Equal(t, arg.CreatedBy, user.CreatedBy)
	require.True(t, util.CheckPasswordHash(user.Password, password))

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestPostUser(t *testing.T) {
	createRandomUser(t)
}
