package service

import (
	"gocard/db"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDrive = "postgres"
	ENV     = "TEST"
)

func TestMain(m *testing.M) {
	db.NewDB(ENV, "../")
	os.Exit(m.Run())
}
