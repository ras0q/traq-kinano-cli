package repository

import "github.com/Ras96/traq-kinano-cli/usecases/repository"

type repositories struct{}

func NewRepositories() repository.Repositories {
	return &repositories{}
}
