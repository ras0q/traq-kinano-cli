package repository_test

import (
	"testing"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ras0q/traq-kinano-cli/ent"
)

type testRepository struct {
	client *ent.Client
	mock   sqlmock.Sqlmock
}

func setupTestRepository(t *testing.T) *testRepository {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(t.Log))

	return &testRepository{
		client: client,
		mock:   mock,
	}
}
