package apitest

import (
	"context"
	"gocard/db"
	"gocard/enum"
	server "gocard/route"
	"gocard/util"
	"os"
	"testing"

	sqlc "gocard/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	r = server.InitRouter()
	db.NewDB("TEST", "../")
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T, password string) sqlc.User {
	if password == "" {
		password = util.RandomPassword()
	}
	arg := sqlc.PostUserParams{
		Name:          util.RandomUser(),
		Email:         util.RandomEmail(),
		Password:      util.HashPassword(password),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}

	user, err := db.Queries.PostUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}
