package mock_repository //nolint:revive,stylecheck

import (
	gomock "github.com/golang/mock/gomock"
	repository "github.com/ras0q/traq-kinano-cli/usecases/repository"
)

type MockRepositories struct {
	*MockAliasRepository
}

func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	return &MockRepositories{
		NewMockAliasRepository(ctrl),
	}
}

var _ repository.Repositories = (*MockRepositories)(nil)
