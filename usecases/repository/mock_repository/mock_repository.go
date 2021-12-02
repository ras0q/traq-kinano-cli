package mock_repository //nolint:revive,stylecheck

import (
	repository "github.com/Ras96/traq-kinano-cli/usecases/repository"
	gomock "github.com/golang/mock/gomock"
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
