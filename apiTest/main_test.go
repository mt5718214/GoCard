package apitest

import (
	"context"
	"gocard/db"
	"gocard/enum"
	server "gocard/route"
	"gocard/util"
	"log"
	"os"
	"testing"

	sqlc "gocard/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	r = server.InitRouter()
	db.NewDB("TEST", "../")
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T, password string, isAdmin bool) sqlc.User {
	var isAdminInt int16
	if isAdmin {
		isAdminInt = 1
	}
	if password == "" {
		password = util.RandomPassword()
	}
	arg := sqlc.PostUserParams{
		Name:          util.RandomUser(),
		Email:         util.RandomEmail(),
		Password:      util.HashPassword(password),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
		IsAdmin:       isAdminInt,
	}

	user, err := db.Queries.PostUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func generateTestToken(userId uuid.UUID, username, email string, isAdmin int16) (token string) {
	token, err := util.CreateJWT("testToken", userId, username, email, isAdmin)
	if err != nil {
		log.Fatal("create token err: ", err.Error())
	}
	return
}
