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

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Password, user2.Password)

	require.NotZero(t, user2.ID)
}
