package db

import (
	"context"
	"gocard/enum"
	"gocard/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomTopic(t *testing.T) uuid.UUID {
	arg := PostTopicParams{
		TopicName:     util.RandomString(5),
		CreatedBy:     enum.Admin.AdminUuid(),
		LastUpdatedBy: enum.Admin.AdminUuid(),
	}

	topic, err := testQueries.PostTopic(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, topic.ID)
	require.Equal(t, arg.TopicName, topic.TopicName)

	return topic.ID
}

func TestPostTopics(t *testing.T) {
	createRandomTopic(t)
}

func TestUpdateTopic(t *testing.T) {
	topicId := createRandomTopic(t)

	arg := UpdateTopicParams{
		TopicName: util.RandomString(5),
		ID:        topicId,
	}

	topic, err := testQueries.UpdateTopic(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, topic.ID)
	require.Equal(t, arg.TopicName, topic.TopicName)
}
