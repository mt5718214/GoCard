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

type UpdateTopicsReq struct {
	TopicName string
}

func TestUpdateTopics(t *testing.T) {
	topicId := generateRandomTopic(t)
	user := createRandomUser(t, util.RandomPassword(), true)
	token := generateTestToken(user.ID, user.Name, user.Email, user.IsAdmin)

	arg := UpdateTopicsReq{
		TopicName: util.RandomString(5),
	}
	jsonValue, err := json.Marshal(arg)
	if err != nil {
		log.Fatal("convert to json error: ", err.Error())
	}

	url := "/dev/api/v1/admin/topics/" + topicId.String()
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}
