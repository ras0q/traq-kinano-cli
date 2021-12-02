//nolint:wsl
package handler_test

import (
	"context"
	"testing"

	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/usecases/repository/mock_repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/Ras96/traq-kinano-cli/util/assert"
	"github.com/Ras96/traq-kinano-cli/util/random"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_handlers_CallAlias(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx   context.Context
		short string
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Alias
		setup   func(r *mock_repository.MockRepositories, s service.Services, args args, want *ent.Alias)
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				short: "test",
			},
			want: &ent.Alias{
				ID:     random.UUID(),
				UserID: random.UUID(),
				Short:  "test",
				Long:   "testtest",
			},
			setup: func(r *mock_repository.MockRepositories, s service.Services, args args, want *ent.Alias) {
				r.EXPECT().CallAlias(args.ctx, args.short).Return(want, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockRepositories(ctrl)
			srv := service.NewServices()
			tt.setup(repo, srv, tt.args, tt.want)
			h := handler.NewHandlers(repo, srv)
			got, err := h.CallAlias(tt.args.ctx, tt.args.short)
			assert.Error(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got, cmpopts.IgnoreUnexported(ent.Alias{}))
		})
	}
}
