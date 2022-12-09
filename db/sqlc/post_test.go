package db

import (
	"context"
	"testing"
	// "github.com/stretchr/testify/require"
)

func TestPostposts(t *testing.T) {
	arg := PostUserParams{}
	user := testQueries.db.PostUser(context.Background())
}
