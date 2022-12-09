package repository

import (
	"github.com/ras0q/traq-kinano-cli/ent"
	"github.com/ras0q/traq-kinano-cli/usecases/repository"
)

type repositories struct {
	ent *ent.Client
}

func NewRepositories(ent *ent.Client) repository.Repositories {
	return &repositories{ent}
}
