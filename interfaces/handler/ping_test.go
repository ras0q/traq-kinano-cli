//nolint:wsl
package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ras0q/traq-kinano-cli/interfaces/handler"
	"github.com/ras0q/traq-kinano-cli/usecases/repository/mock_repository"
	"github.com/ras0q/traq-kinano-cli/usecases/service"
	"github.com/ras0q/traq-kinano-cli/util/assert"
)

func Test_handlers_Ping(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		setup   func(r *mock_repository.MockRepositories, s service.Services)
		wantErr bool
	}{
		{
			name:    "Success",
			setup:   func(r *mock_repository.MockRepositories, s service.Services) {},
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
			tt.setup(repo, srv)
			h := handler.NewHandlers(repo, srv)
			assert.Error(t, tt.wantErr, h.Ping())
		})
	}
}
