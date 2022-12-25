package apitest

import (
	"bytes"
	"context"
	"encoding/json"
	authCtrl "gocard/controllers"
	"gocard/util"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	db "gocard/db"

	"github.com/stretchr/testify/require"
)

type resMessage struct {
	Message, Data string
}

func TestRegisterHandler(t *testing.T) {
	password := util.RandomPassword()
	arg := authCtrl.SignupReqBody{
		Name:          util.RandomUser(),
		Email:         util.RandomEmail(),
		Password:      password,
		CheckPassword: password,
	}
	// convert arg struct to JSON-encoded data
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	req := httptest.NewRequest("POST", "/dev/api/v1/signup", bytes.NewBuffer(jsonValue)) // 建立一個請求
	w := httptest.NewRecorder()                                                          // 建立一個ResponseRecorder其實作http.ResponseWriter，用來記錄response狀態
	r.ServeHTTP(w, req)                                                                  // gin.Engine.ServerHttp實作http.Handler介面，用來處理HTTP請求及回應。

	var res resMessage
	json.Unmarshal(w.Body.Bytes(), &res)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, "Create user success", res.Message)

	user, err := db.Queries.GetUserByEmail(context.Background(), arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
}

func TestAuthHandler(t *testing.T) {
	password := util.RandomPassword()
	user := createRandomUser(t, password, false)
	arg := authCtrl.LoginReqBody{
		Email:    user.Email,
		Password: password,
	}
	// convert arg struct to JSON-encoded data
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	req := httptest.NewRequest("POST", "/dev/api/v1/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var res resMessage
	json.Unmarshal(w.Body.Bytes(), &res)

	require.Equal(t, http.StatusOK, w.Code)
	require.NotEmpty(t, res.Data)
}
