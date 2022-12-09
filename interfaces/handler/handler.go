package handler

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/ras0q/traq-kinano-cli/ent"
	"github.com/ras0q/traq-kinano-cli/usecases/repository"
	"github.com/ras0q/traq-kinano-cli/usecases/service"
)

type Handlers interface {
	// Alias
	CallAlias(ctx context.Context, short string) (*ent.Alias, error)
	AddAlias(ctx context.Context, userID uuid.UUID, short string, long string) error
	// Ping
	Ping() error
}

type handlers struct {
	Repo repository.Repositories
	Srv  service.Services
}

func NewHandlers(repo repository.Repositories, srv service.Services) Handlers {
	return &handlers{
		Repo: repo,
		Srv:  srv,
	}
}
