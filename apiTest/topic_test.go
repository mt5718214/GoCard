package apitest

import (
	"bytes"
	"context"
	"encoding/json"
	"gocard/db"
	sqlc "gocard/db/sqlc"
	"gocard/enum"
	"gocard/util"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func generateRandomTopic(t *testing.T) uuid.UUID {
	arg := sqlc.PostTopicParams{
		TopicName:     util.RandomString(5),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}
	topic, err := db.Queries.PostTopic(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, topic.ID)
	require.Equal(t, arg.TopicName, topic.TopicName)

	return topic.ID
}

type postTopicReq struct {
	TopicName string
}

func TestPostTopic(t *testing.T) {
	user := createRandomUser(t, util.RandomPassword(), true)
	token := generateTestToken(user.ID, user.Name, user.Email, user.IsAdmin)

	arg := postTopicReq{
		TopicName: util.RandomString(5),
	}
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	req, err := http.NewRequest("POST", "/dev/api/v1/admin/topics/", bytes.NewBuffer(jsonValue))
	require.NoError(t, err)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
}

type updateTopicReq struct {
	TopicName string
}

func TestUpdateTopic(t *testing.T) {
	topicId := generateRandomTopic(t)
	user := createRandomUser(t, util.RandomPassword(), true)
	token := generateTestToken(user.ID, user.Name, user.Email, user.IsAdmin)

	arg := updateTopicReq{
		TopicName: util.RandomString(5),
	}
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	url := "/dev/api/v1/admin/topics/" + topicId.String()
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	require.NoError(t, err)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteTopic(t *testing.T) {
	topicID := generateRandomTopic(t)
	user := createRandomUser(t, util.RandomPassword(), true)
	token := generateTestToken(user.ID, user.Name, user.Email, user.IsAdmin)

	url := "/dev/api/v1/admin/topics/" + topicID.String()
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}
