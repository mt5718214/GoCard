package apitest

import (
	"bytes"
	"encoding/json"
	"gocard/util"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type postTopicsReq struct {
	TopicName string
}

func TestPostTopics(t *testing.T) {
	user := createRandomUser(t, util.RandomPassword(), true)
	token := generateTestToken(user.ID, user.Name, user.Email, user.IsAdmin)

	arg := postTopicsReq{
		TopicName: util.RandomString(5),
	}
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	req, _ := http.NewRequest("POST", "/dev/api/v1/admin/topics/", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
}
