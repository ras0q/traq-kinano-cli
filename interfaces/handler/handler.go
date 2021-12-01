package handler

import (
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/gofrs/uuid"
)

type Handlers interface {
	// Alias
	CallAlias(channelID string, short string) error
	AddAlias(channelID string, userID uuid.UUID, short string, long string) error
	// Ping
	Ping(channelID string) error
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
