package repository_test

import (
	"testing"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/google/go-cmp/cmp"
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

func assertErr(t *testing.T, wantErr bool, err error) {
	t.Helper()

	if !wantErr && err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func assertEqual(t *testing.T, want interface{}, got interface{}, opts ...cmp.Option) {
	t.Helper()

	if diff := cmp.Diff(want, got, opts...); len(diff) > 0 {
		t.Error(diff)
	}
}
