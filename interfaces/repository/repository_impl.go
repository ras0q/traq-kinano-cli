package repository

import (
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
)

type repositories struct {
	ent *ent.Client
}

func NewRepositories(ent *ent.Client) repository.Repositories {
	return &repositories{ent}
}
