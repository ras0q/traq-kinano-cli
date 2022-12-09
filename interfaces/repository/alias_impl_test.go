//nolint:wsl,funlen
package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ras0q/traq-kinano-cli/ent"
	impl "github.com/ras0q/traq-kinano-cli/interfaces/repository"
	repo "github.com/ras0q/traq-kinano-cli/usecases/repository"
	"github.com/ras0q/traq-kinano-cli/util/assert"
	"github.com/ras0q/traq-kinano-cli/util/random"
)

func Test_repositories_CallAlias(t *testing.T) {
	t.Parallel()
	type args struct {
		short string
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Alias
		setup   func(tr *testRepository, args args, want *ent.Alias)
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				short: "test",
			},
			want: &ent.Alias{
				ID:     random.UUID(),
				UserID: random.UUID(),
				Short:  "test",
				Long:   "testtest",
			},
			setup: func(tr *testRepository, args args, want *ent.Alias) {
				tr.mock.ExpectQuery("SELECT").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "user_id", "short", "long"}).
							AddRow(want.ID, want.UserID, want.Short, want.Long),
					)
			},
			wantErr: false,
		},
		{
			name: "ErrFirst",
			args: args{
				short: "test",
			},
			want: nil,
			setup: func(tr *testRepository, args args, want *ent.Alias) {
				tr.mock.ExpectQuery("SELECT").
					WillReturnError(errors.New("error getting first"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := setupTestRepository(t)
			defer tr.client.Close()
			r := impl.NewRepositories(tr.client)
			tt.setup(tr, tt.args, tt.want)
			got, err := r.CallAlias(context.Background(), tt.args.short)
			assert.Error(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got, cmpopts.IgnoreUnexported(ent.Alias{}))
		})
	}
}

func Test_repositories_AddAlias(t *testing.T) {
	uid := random.UUID()

	t.Parallel()
	type args struct {
		args *repo.AddAliasArgs
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Alias
		setup   func(tr *testRepository, args args, want *ent.Alias)
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				args: &repo.AddAliasArgs{
					UserID: uid,
					Short:  "test",
					Long:   "testtest",
				},
			},
			want: &ent.Alias{
				ID:     uuid.UUID{}, // do not compare
				UserID: uid,
				Short:  "test",
				Long:   "testtest",
			},
			setup: func(tr *testRepository, args args, want *ent.Alias) {
				tr.mock.ExpectBegin()
				tr.mock.ExpectExec("INSERT INTO").
					WithArgs(args.args.UserID, args.args.Short, args.args.Long, sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				tr.mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "ErrCreate",
			args: args{
				args: &repo.AddAliasArgs{
					UserID: random.UUID(),
					Short:  "test",
					Long:   "testtest",
				},
			},
			want: nil,
			setup: func(tr *testRepository, args args, want *ent.Alias) {
				tr.mock.ExpectBegin()
				tr.mock.ExpectExec("INSERT INTO").
					WithArgs(args.args.UserID, args.args.Short, args.args.Long, sqlmock.AnyArg()).
					WillReturnError(errors.New("error creating"))
				tr.mock.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := setupTestRepository(t)
			defer tr.client.Close()
			r := impl.NewRepositories(tr.client)
			tt.setup(tr, tt.args, tt.want)
			got, err := r.AddAlias(context.Background(), tt.args.args)
			assert.Error(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got, cmpopts.IgnoreUnexported(ent.Alias{}), cmpopts.IgnoreFields(ent.Alias{}, "ID"))
		})
	}
}
